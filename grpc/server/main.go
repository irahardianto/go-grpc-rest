package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
)

const (
	port = ":18288"
)

type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResult, error) {
	return &pb.ScoreResult{Score: "Hello this is score of " + in.GetPlayer1() + " & " + in.GetPlayer2()}, nil
}

func main() {

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterScoreBoardServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
