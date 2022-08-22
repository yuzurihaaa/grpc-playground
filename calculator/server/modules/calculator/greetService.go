package calculator

import (
	"context"
	pb "grpc-playground/calculator/proto"
	"log"
	"time"
)

type Service struct {
	pb.UnsafeCalculatorServiceServer
}

func (s *Service) Sum(_ context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum invoked with %v\n", in)
	return &pb.SumResponse{Total: in.FirstNumber + in.SecondNumber}, nil
}

func (s *Service) Primes(in *pb.Number, stream pb.CalculatorService_PrimesServer) error {
	divisor := int64(2)
	N := in.Number
	for N > 1 {
		if N%divisor == 0 {
			time.Sleep(time.Second)
			err := stream.Send(&pb.Number{Number: divisor})
			if err != nil {
				return err
			}
			N /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
