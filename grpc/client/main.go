package main

import (
	"context"
	"log"

	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:18288"
	player1Name = "Rafael"
	player2Name = "Serena"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewScoreBoardClient(conn)

	result, err := client.GetScore(context.Background(), &pb.ScoreRequest{Player1: player1Name, Player2: player2Name})
	if err != nil {
		log.Fatalf("could not get score: %v", err)
	}
	log.Printf("The current score is: %s", result.GetScore())
}
