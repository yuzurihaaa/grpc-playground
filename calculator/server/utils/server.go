package utils

import (
	"google.golang.org/grpc"
	pb "grpc-playground/calculator/proto"
	"grpc-playground/calculator/server/modules/calculator"
	"log"
	"net"
)

func RegisterServer(lis net.Listener) {
	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &calculator.Service{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
