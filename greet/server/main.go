package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc-playground/greet/proto"
)

const addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen to %v", err)
	}
	log.Printf("Listening on %v", addr)

	server := grpc.NewServer()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}

type Server struct {
	pb.GreetServiceServer
}
