package greet

import (
	"context"
	"fmt"
	pb "grpc-playground/greet/proto"
	"io"
	"log"
	"time"
)

type Service struct {
	pb.UnimplementedGreetServiceServer
}

func (s *Service) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}

func (s *Service) Greets(in *pb.GreetRequest, stream pb.GreetService_GreetsServer) error {
	log.Printf("Greets invoked with %v\n", in)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		res := fmt.Sprintf("Hello %v, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

func (s *Service) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Print("LongGreet invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			err := stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
			if err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			log.Fatalf("Fail to receive data %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
	return nil
}
