package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/awesomeProject/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, in *pb.Input) (*pb.Result, error) {
	log.Printf("Received: %d, %d", in.GetA(), in.GetB())
	return &pb.Result{Res: in.GetA() + in.GetB()}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.Input) (*pb.Result, error) {
	log.Printf("Received: %d, %d", in.GetA(), in.GetB())
	return &pb.Result{Res: in.GetA() - in.GetB()}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.Input) (*pb.Result, error) {
	log.Printf("Received: %d, %d", in.GetA(), in.GetB())
	if in.GetB() == 0 {
		return &pb.Result{Res: 0}, nil
	}
	return &pb.Result{Res: in.GetA() / in.GetB()}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.Input) (*pb.Result, error) {
	log.Printf("Received: %d, %d", in.GetA(), in.GetB())
	return &pb.Result{Res: in.GetA() * in.GetB()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
