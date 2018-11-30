package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("server:SayHello")

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
