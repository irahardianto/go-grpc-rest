package servers

import (
	"context"
	"testing"

	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
)

func TestGetScore(t *testing.T) {
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "Rafael", "Serena").Return("Forty-Fifteen", nil)

	scoreBoardServer := ScoreBoardServer{}
	scoreBoardServer.PlayerService = playerService

	req := &pb.ScoreRequest{Player1: "Rafael", Player2: "Serena"}
	resp, err := scoreBoardServer.GetScore(context.Background(), req)
	if err != nil {
		t.Errorf("HelloTest(%v) got unexpected error", err)
	}
	if resp.GetScore() != "Forty-Fifteen" {
		t.Errorf("Get=%v, wanted %v", resp.GetScore(), "Forty-Fifteen")
	}
}
