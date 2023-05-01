package services

import (
	"fmt"

	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/repositories"
)

type playerService struct {
	playerRepository interfaces.IPlayerRepository
}

func NewPlayerService() interfaces.IPlayeService {
	return playerService{
		playerRepository: repositories.NewPlayerRepository(),
	}
}

func (ps playerService) CreatePlayer(player models.Player) models.Player {
	fmt.Printf("\nCreate a player %s ...", player.Name)
	playerCreated := ps.playerRepository.Save(player)
	return playerCreated
}

func (ps playerService) GetAll() []models.Player {
	return ps.playerRepository.FindAll()
}

func (ps playerService) GetPlayerById(id int16) models.Player {
	return ps.playerRepository.FindById(id)
}
