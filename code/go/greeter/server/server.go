package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "server/greeter"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRes, error) {
	fmt.Println("server:SayHello")
	return &pb.HelloRes{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHellos(in *pb.HelloReq, stream pb.Greeter_SayHellosServer) error {
	fmt.Println("server:SayHellos")

	done := make(chan bool)
	defer close(done)

	errCh := make(chan error)
	defer close(errCh)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	go func(name string, max int32) {
		var n int32
		for range ticker.C {
			if n < max {
				msg := &pb.HelloRes{Message: "Hello " + in.Name}
				if err := stream.Send(msg); err != nil {
					errCh <- err
					done <- true
				}
				n++
			} else {
				errCh <- nil
				done <- true
			}
		}
	}(in.Name, in.Count)

	err, _ := <-errCh, <-done
	return err
}

func (s *server) GreetMany(stream pb.Greeter_GreetManyServer) error {
	fmt.Println("server:GreetMany")

	names := make([]string, 0, 5)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			msg := &pb.HelloRes{Message: "Hello " + strings.Join(names, ", ")}
			return stream.SendAndClose(msg)
		}
		if err != nil {
			return err
		}
		names = append(names, in.Name)
	}
}

func (s *server) GreetChat(stream pb.Greeter_GreetChatServer) error {
	fmt.Println("server:GreetChat")

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		msg := &pb.HelloRes{Message: "Hello " + in.Name}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
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
