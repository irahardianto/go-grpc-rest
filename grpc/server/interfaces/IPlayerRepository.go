package interfaces

import (
	"github.com/irahardianto/go-grpc-rest/grpc/server/models"
)

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
