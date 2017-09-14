package main

import (
	"sync"

	"github.com/irahardianto/go-grpc-rest/grpc/server/infrastructures"
	"github.com/irahardianto/go-grpc-rest/grpc/server/repositories"
	"github.com/irahardianto/go-grpc-rest/grpc/server/servers"
	"github.com/irahardianto/go-grpc-rest/grpc/server/services"
)

type IServiceContainer interface {
	InjectScoreBoardServer() servers.ScoreBoardServer
}

type kernel struct{}

func (k *kernel) InjectScoreBoardServer() servers.ScoreBoardServer {

	sqlconn := new(infrastructures.SqlConnection)
	sqlconn.InitDB()

	playerRepository := new(repositories.PlayerRepository)
	playerRepository.Db.Db = sqlconn.GetDB()

	playerService := new(services.PlayerService)
	playerService.PlayerRepository = playerRepository

	scoreBoardServer := servers.ScoreBoardServer{}
	scoreBoardServer.PlayerService = playerService

	return scoreBoardServer
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
