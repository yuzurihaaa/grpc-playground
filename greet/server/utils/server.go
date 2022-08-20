package utils

import (
	"google.golang.org/grpc"
	pb "grpc-playground/greet/proto"
	"grpc-playground/greet/server/modules/greet"
	"log"
	"net"
)

func RegisterServer(lis net.Listener) {
	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server, &greet.Service{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
