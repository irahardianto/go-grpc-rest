package servers

import (
	"context"

	"github.com/irahardianto/go-grpc-rest/grpc/server/interfaces"
	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
)

type ScoreBoardServer struct {
	PlayerService interfaces.IPlayerService
}

func (server *ScoreBoardServer) GetScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResult, error) {

	scores, err := server.PlayerService.GetScores(in.GetPlayer1(), in.GetPlayer2())
	if err != nil {
		//Handle error
	}

	return &pb.ScoreResult{Score: scores}, nil
}
