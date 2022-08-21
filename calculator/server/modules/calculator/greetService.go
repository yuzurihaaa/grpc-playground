package calculator

import (
	"context"
	pb "grpc-playground/calculator/proto"
	"log"
)

type Service struct {
	pb.UnsafeCalculatorServiceServer
}

func (s *Service) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum invoked with %v\n", in)
	return &pb.SumResponse{Total: in.FirstNumber + in.SecondNumber}, nil
}
