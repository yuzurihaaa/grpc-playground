package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-playground/greet/proto"
	"io"
	"log"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	// Context parsing is overkill, but I'm just playing around.
	greetClient := pb.NewGreetServiceClient(conn)
	doGreet(greetClient)
	listenGreet(greetClient)
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
