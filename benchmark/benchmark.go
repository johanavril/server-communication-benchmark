package main

import (
	"github.com/johanavril/server-communication-benchmark/pb"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

const (
	restAddr = "http://localhost:8081"
	grpcAddr = "localhost:8082"
)

type server struct {
	rest *http.Client
	grpc pb.BenchServiceClient

	grpcConn *grpc.ClientConn
}

func New() (*server, error) {
	rest := &http.Client{Timeout: 30 * time.Second}
	cc, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	grpc := pb.NewBenchServiceClient(cc)


	return &server{
		rest: rest,
		grpc: grpc,
		grpcConn: cc,
	}, nil
}

func (s *server) TearDown() {
	s.grpcConn.Close()
}
