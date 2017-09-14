package servers

import (
	"context"

	"github.com/irahardianto/go-grpc-rest/grpc/server/interfaces"
	pb "github.com/irahardianto/go-grpc-rest/grpc/tennisscore"
)

type ScoreBoardServer struct {
	PlayerService interfaces.IPlayerService
}

func (s *ScoreBoardServer) GetScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResult, error) {
	return &pb.ScoreResult{Score: "Hello this is score of " + in.GetPlayer1() + " & " + in.GetPlayer2()}, nil
}
