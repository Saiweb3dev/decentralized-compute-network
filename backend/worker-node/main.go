package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto"

	"google.golang.org/grpc"
)

//Server struct implements HelloService
type Server struct {
	pb.UnimplementedHelloServiceServer
}

//SayHelloBro impletements the gRPC method

func (s *Server) SayHelloBro(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Received message from client:",req.Message)
	return &pb.HelloResponse{Reply: "Hi from Go!"}, nil
}

func main() {
	//Start listening on port 5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})

	fmt.Println("gRPC Server is running on port 5001...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}