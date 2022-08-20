package main

import (
	"grpc-playground/greet/server/utils"
	"log"
	"net"
)

const addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen to %v", err)
	}
	log.Printf("Listening on %v", addr)

	utils.RegisterServer(lis)
}
