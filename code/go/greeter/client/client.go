package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "client/greeter"

	"google.golang.org/grpc"
)

func sayHello(client pb.GreeterClient) {
	fmt.Println("client:sayHello")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloReq{Name: "world"})
	if err != nil {
		log.Fatalf("SayHello error: %v", err)
	}

	log.Printf("Greeting: %s", res.Message)
}

func sayHellos(client pb.GreeterClient) {
	fmt.Println("client:sayHellos")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHellos(ctx, &pb.HelloReq{Name: "world", Count: 5})
	if err != nil {
		log.Fatalf("SayHellos error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.SayHellos(_) = _, %v", client, err)
		}
		log.Println(res.Message)
	}
}

func greetMany(client pb.GreeterClient) {
	fmt.Println("client:greetMany")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GreetMany(ctx)
	if err != nil {
		log.Fatalf("GreetMany error: %v", err)
	}

	done := make(chan bool)
	defer close(done)

	errCh := make(chan error)
	defer close(errCh)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	NAMES := [4]string{"Bob", "Kate", "Jim", "Sara"}

	go func(names []string) {
		var n int
		for range ticker.C {
			if n < len(names) {
				msg := &pb.HelloReq{Name: names[n]}
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
	}(NAMES[:])

	err, _ = <-errCh, <-done
	if err != nil {
		log.Fatalf("%v.GreetMany(_) = _, %v", client, err)
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	log.Printf("Greeting: %v", reply.Message)
}

func greetChat(client pb.GreeterClient) {
	fmt.Println("client:greetChat")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GreetChat(ctx)
	if err != nil {
		log.Fatalf("GreetMany error: %v", err)
	}

	// read side
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a greeting: %v", err)
			}
			log.Printf("Greeting: %+v", in.Message)
		}
	}()

	// send side
	done := make(chan bool)
	defer close(done)

	errCh := make(chan error)
	defer close(errCh)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	NAMES := [4]string{"Bob", "Kate", "Jim", "Sara"}

	go func(names []string) {
		var n int
		for range ticker.C {
			if n < len(names) {
				msg := &pb.HelloReq{Name: names[n]}
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
	}(NAMES[:])

	_, _ = <-errCh, <-done

	stream.CloseSend()

	<-waitc
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	sayHello(c)

	sayHellos(c)

	greetMany(c)

	greetChat(c)
}
