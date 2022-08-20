package greet

import (
	"context"
	pb "grpc-playground/greet/proto"
	"log"
)

type Service struct {
	pb.UnimplementedGreetServiceServer
}

func (s *Service) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
