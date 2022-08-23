package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-playground/greet/proto"
	"io"
	"log"
	"time"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	greetClient := pb.NewGreetServiceClient(conn)
	doGreet(greetClient)
	listenGreet(greetClient)
	sendGreets(greetClient)
}

func doGreet(client pb.GreetServiceClient) {
	log.Println("modules.greet - DoGreet")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Yusuf",
	})

	if err != nil {
		log.Fatalf("Fail to greet server %v\n", err)
	}

	log.Printf("Greeting response %v\n", res.Result)
}

func listenGreet(client pb.GreetServiceClient) {
	log.Println("modules.greet - DoGreet")

	stream, err := client.Greets(context.Background(), &pb.GreetRequest{
		FirstName: "Yusuf",
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

		log.Printf("listenGreet response %v\n", message.Result)
	}

}

func sendGreets(client pb.GreetServiceClient) {
	log.Println("modules.greet - DoGreet")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kamal"},
		{FirstName: "Firdaus"},
		{FirstName: "Hafiz"},
	}

	stream, err := client.LongGreet(context.Background())

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

	log.Printf("Greet %s", res.Result)
}
