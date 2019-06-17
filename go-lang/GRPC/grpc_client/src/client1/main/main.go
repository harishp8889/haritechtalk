package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "client1/main/pb"
)

const (
	address = "localhost:17009"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRpcCommunicator1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetRpc1Msg(ctx, &pb.RequestParams{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.Rpc1Msg)
}

