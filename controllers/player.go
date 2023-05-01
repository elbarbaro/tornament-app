package controllers

import (
	"fmt"

	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/services"
)

type playerController struct {
	playerService interfaces.IPlayeService
}

func NewPlayerController() playerController {
	return playerController{
		playerService: services.NewPlayerService(),
	}
}

func (pc playerController) CreatePlayer(name string, number int8, country string, position string, role string) models.Player {
	player := models.Player{
		Name:     name,
		Number:   number,
		Country:  country,
		Position: position,
		Role:     role,
	}
	playerCreated := pc.playerService.CreatePlayer(player)

	fmt.Printf("\nNew player %s successfully \n", player.Name)

	return playerCreated
}

func (pc playerController) GetAll() []models.Player {
	return pc.playerService.GetAll()
}

func (pc playerController) GetById(id int16) models.Player {
	return pc.playerService.GetPlayerById(id)
}
