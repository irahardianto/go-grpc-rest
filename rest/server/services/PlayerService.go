package services

import (
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/models"
)

type PlayerService struct {
	PlayerRepository interfaces.IPlayerRepository
}

func (service *PlayerService) GetScores(player1Name string, player2Name string) (string, error) {

	baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
	var result string

	player1, err := service.PlayerRepository.GetPlayerByName(player1Name)
	if err != nil {
		//Handle error
	}

	player2, err := service.PlayerRepository.GetPlayerByName(player2Name)
	if err != nil {
		//Handle error
	}

	if player1.Score < 4 && player2.Score < 4 && !(player1.Score+player2.Score == 6) {

		s := baseScore[player1.Score]

		if player1.Score == player2.Score {
			result = s + "-All"
		} else {
			result = s + "-" + baseScore[player2.Score]
		}
	}

	if player1.Score == player2.Score {
		result = "Deuce"
	}

	return result, nil
}

func (service *PlayerService) GetPlayerMessage() models.MessageModel {

	data := service.PlayerRepository.GetPlayerMessageFromAPI()

	return data
}
