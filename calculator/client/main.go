package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-playground/calculator/proto"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const addr = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	calculatorClient := pb.NewCalculatorServiceClient(conn)
	if len(os.Args[1:]) == 2 {
		first, err := strconv.ParseFloat(os.Args[1:][0], 32)
		if err != nil {
			panic(err)
		}
		second, err := strconv.ParseFloat(os.Args[1:][1], 32)
		if err != nil {
			panic(err)
		}

		doSum(calculatorClient, first, second)
		return
	}

	if len(os.Args[1:]) == 1 {
		first, err := strconv.ParseInt(os.Args[1:][0], 10, 64)
		if err != nil {
			panic(err)
		}
		listenToPrime(calculatorClient, first)
		return
	}

	sendNumbers(calculatorClient)
	CurrentMax(calculatorClient)
}

func doSum(client pb.CalculatorServiceClient, x, y float64) {
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  x,
		SecondNumber: y,
	})

	if err != nil {
		log.Fatalf("Fail to greet server %v\n", err)
	}

	log.Printf("Greeting response %v\n", res.Total)
}

func listenToPrime(client pb.CalculatorServiceClient, x int64) {
	stream, err := client.Primes(context.Background(), &pb.Number{
		Number: x,
	})

	if err != nil {
		log.Fatalf("Fail to greet server %v\n", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Fail to receive data %v\n", err)
		}

		log.Printf("prime number response %v\n", message.Number)
	}
}

func sendNumbers(client pb.CalculatorServiceClient) {
	log.Println("modules.greet - DoGreet")

	reqs := []*pb.NumberFloat{
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("Fail to greet server %v\n", err)
	}

	for _, req := range reqs {

		log.Printf("sending %v", req)

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Fail to send data %v\n", err)
		}

		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Fail to CloseAndRecv %v\n", err)
	}

	log.Printf("Greet %v", res.Number)
}

func CurrentMax(client pb.CalculatorServiceClient) {
	log.Println("greetings invoke")

	stream, err := client.CurrentMax(context.Background())

	if err != nil {
		log.Fatalf("Error creating stream %v\n", err)
	}

	reqs := []*pb.Number{
		{Number: 20},
		{Number: 50},
		{Number: 40},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {

			log.Printf("sending %v", req)

			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Fail to send data %v\n", err)
			}

			time.Sleep(time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Fail to CloseSend %v\n", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Fail to receive data %v\n", err)
				break
			}

			log.Printf("Current max %v", res.Number)
		}

		close(waitc)
	}()

	<-waitc
}
