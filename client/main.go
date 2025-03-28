package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pascalallen/grpc-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:8008: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloWorldRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	log.Printf("resoposnse recieved: %v", r.GetMessage())
}
