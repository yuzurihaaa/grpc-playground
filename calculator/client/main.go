package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-playground/calculator/proto"
	"log"
	"os"
	"strconv"
)

const addr = "localhost:50051"

func main() {

	fmt.Println(os.Args[1:][0])
	first, err := strconv.ParseFloat(os.Args[1:][0], 32)
	if err != nil {
		panic(err)
	}
	second, err := strconv.ParseFloat(os.Args[1:][1], 32)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	greetClient := pb.NewCalculatorServiceClient(conn)
	doSum(greetClient, first, second)
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
