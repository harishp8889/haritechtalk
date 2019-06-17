package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"

	log "github.com/sirupsen/logrus"
	pb "rpc2/main/pb"
	pbRpc1 "rpc2/main/pb/rpc1"
)

const (
	servicePort = 17007
	rpc1Address = "localhost:17009"
      )

type server struct{}

func (s *server) GetRpc2Msg(ctx context.Context, in *pb.RequestParams) (*pb.Rpc2Reply, error) {

	return &pb.Rpc2Reply{ Rpc2Msg: "Message From GRPC2-->COMMUNICATION SUCCESS!" }, nil
}

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)

	//De-register the service on exit
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM) //Notify when app is terminated
	signal.Notify(signals, syscall.SIGINT)  //Notify when app is interrupted

	// Shutdown channel to notify the running go-routines to stop when main exists
	shutdownChan := make(chan bool)

	go ShutdownHook(signals, shutdownChan)

	log.Info("RPC2 Started")
	getMessageFromRpc1()

	laddr := fmt.Sprintf(":%d", servicePort)
	lis, lisErr := net.Listen("tcp", laddr)
	if lisErr!= nil {
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterRpcCommunicator2Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getMessageFromRpc1(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(rpc1Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbRpc1.NewRpcCommunicator1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetRpc1Msg(ctx, &pbRpc1.RequestParams{})
	if err != nil {
		log.Fatalf("could not communicate: %v", err)
	}
	log.Info(r.Rpc1Msg)
}

// ShutdownHook is called when service is stopped or aborted based on the OS signals
func ShutdownHook(signals <-chan os.Signal, shutdownChan chan bool) {
	select {
	case <-signals:
		//Go the OS interrupt signal
		log.Println("Deregistering the service")

		//Stop all the routines listening to shutdown channel
		close(shutdownChan)

		os.Exit(3)
	}
}

