package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
)

const (
	port = ":18288"
)

func main() {

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ScoreBoardServer := ServiceContainer().InjectScoreBoardServer()

	pb.RegisterScoreBoardServer(s, &ScoreBoardServer)

	reflection.Register(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
