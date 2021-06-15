package main

import (
	"github.com/johanavril/server-communication-benchmark/pb"
	"google.golang.org/grpc"
	"net"
)

func run() error {
	lis, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterBenchServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
