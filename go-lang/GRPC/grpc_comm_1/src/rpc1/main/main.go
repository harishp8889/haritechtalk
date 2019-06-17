package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	pb "rpc1/main/pb"
)

const servicePort = 17009

type server struct{}

func (s *server) GetRpc1Msg(ctx context.Context, in *pb.RequestParams) (*pb.Rpc1Reply, error) {

	return &pb.Rpc1Reply{ Rpc1Msg: "Message From GRPC1-->COMMUNICATION SUCCESS!" }, nil
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

	log.Info("GRPC1 started!")

	laddr := fmt.Sprintf(":%d", servicePort)
	lis, lisErr := net.Listen("tcp", laddr)
	if lisErr!= nil {
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterRpcCommunicator1Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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

