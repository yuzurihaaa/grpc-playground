package calculator

import (
	"context"
	pb "grpc-playground/calculator/proto"
	"io"
	"log"
	"math"
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

func (s *Service) Average(stream pb.CalculatorService_AverageServer) error {
	var res []float64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			output := float64(0)
			for _, num := range res {
				output += num
			}

			err := stream.SendAndClose(&pb.NumberFloat{
				Number: output / float64(len(res)),
			})
			if err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			log.Fatalf("Fail to receive data %v\n", err)
		}

		res = append(res, req.Number)
	}
	return nil
}

func (s *Service) CurrentMax(stream pb.CalculatorService_CurrentMaxServer) error {
	log.Print("CurrentMax invoked")
	currentMax := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Fail to receive data %v\n", err)
		}

		currentMax = int(int64(math.Max(float64(req.Number), float64(currentMax))))
		err = stream.Send(&pb.Number{
			Number: int64(currentMax),
		})

		if err != nil {
			log.Fatalf("Fail to receive data %v\n", err)
		}
	}
	return nil
}
