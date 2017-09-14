package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/irahardianto/go-grpc-rest/grpc/server/infrastructures"
	"github.com/irahardianto/go-grpc-rest/grpc/server/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepository struct {
	Db infrastructures.SqlConnection
}

func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

	conn := repository.Db.GetDB()

	output := make(chan models.PlayerModel, 1)
	errors := hystrix.Go("get_player_by_name", func() error {

		player := models.PlayerModel{}
		conn.First(&player, "Name = ?", name)
		output <- player
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.PlayerModel{}, err
	}
}
