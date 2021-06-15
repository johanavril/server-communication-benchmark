package main

import (
	"context"
	"fmt"
	"github.com/johanavril/server-communication-benchmark/pb"
)

type server struct {
	pb.UnimplementedBenchServiceServer

	grpc pb.BenchServiceServer
}

func (s *server) Ping(_ context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Pong: "PONG"}, nil
}

func (s *server) Small(_ context.Context, req *pb.SmallRequest) (*pb.SmallResponse, error) {
	username := req.GetIdentity().GetUsername()
	email := req.GetIdentity().GetEmail()
	country := req.GetIdentity().GetCountry()
	return &pb.SmallResponse{
		Summary: fmt.Sprintf("username=%s email=%s country=%s", username, email, country),
	}, nil
}

func (s *server) Big(_ context.Context, req *pb.BigRequest) (*pb.BigResponse, error) {
	username := req.GetIdentity().GetUsername()
	email := req.GetIdentity().GetEmail()
	country := req.GetIdentity().GetCountry()
	lat := req.GetLocation().GetLat()
	lon := req.GetLocation().GetLon()
	interest := req.GetInterest()
	bookmark := req.GetBookmark()

	bookmarkedInterest := make(map[string]bool, len(interest))
	organizedBookmark := make(map[string]*pb.ContentList)

	for _, bm := range bookmark {
		var cl *pb.ContentList
		if v, ok := organizedBookmark[bm.GetGenre()]; !ok {
			cl = &pb.ContentList{}
		} else {
			cl = v
		}

		cl.Content = append(cl.GetContent(), bm)
		organizedBookmark[bm.GetGenre()] = cl
	}

	for _, v := range interest {
		_, ok := organizedBookmark[v]
		bookmarkedInterest[v] = ok
	}

	return &pb.BigResponse{
		Summary: fmt.Sprintf("username=%s email=%s country=%s lat=%f lon=%f", username, email, country, lat, lon),
		BookmarkedInterest: bookmarkedInterest,
		OrganizedBookmark: organizedBookmark,
	}, nil
}

