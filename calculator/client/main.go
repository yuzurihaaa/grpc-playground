package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-playground/calculator/proto"
	"io"
	"log"
	"os"
	"strconv"
)

const addr = "localhost:50051"

func main() {

	fmt.Println(os.Args[1:][0])
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	greetClient := pb.NewCalculatorServiceClient(conn)
	if len(os.Args[1:]) == 2 {
		first, err := strconv.ParseFloat(os.Args[1:][0], 32)
		if err != nil {
			panic(err)
		}
		second, err := strconv.ParseFloat(os.Args[1:][1], 32)
		if err != nil {
			panic(err)
		}

		doSum(greetClient, first, second)
	}

	if len(os.Args[1:]) == 1 {
		first, err := strconv.ParseInt(os.Args[1:][0], 10, 64)
		if err != nil {
			panic(err)
		}
		listenToPrime(greetClient, first)
	}
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

		log.Printf("listenGreet response %v\n", message.Number)
	}
}
