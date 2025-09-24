package main

import (
	"context"
	"grps/pb/grps/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)



type server struct {
	pb.UnimplementedHelloServiceServer
}


func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}


func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	log.Println("grps server running  on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}