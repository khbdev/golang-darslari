package main

import (
	"context"
	"grps/pb/grps/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Serverga ulanamiz
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Clientni yaratamiz
	c := pb.NewHelloServiceClient(conn)

	// Timeout bilan context yaratamiz
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// RPC chaqiramiz
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Azizbek"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", res.Message)
}
