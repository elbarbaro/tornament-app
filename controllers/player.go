package controllers

import (
	"fmt"

	"elbarbaro.com/tornament-app/handlers/requests"
	"elbarbaro.com/tornament-app/handlers/responses"
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/services"
)

type PlayerController struct {
	playerService interfaces.IPlayeService
}

func NewPlayerController() PlayerController {
	return PlayerController{
		playerService: services.NewPlayerService(),
	}
}

func (pc PlayerController) CreatePlayer(playerRequest requests.PlayerRequest) responses.PlayerResponse {
	player := models.Player{
		Name:     playerRequest.Name,
		Number:   playerRequest.Number,
		Country:  playerRequest.Country,
		Position: playerRequest.Position,
		Role:     playerRequest.Role,
	}
	playerCreated := pc.playerService.CreatePlayer(player)

	fmt.Printf("\nNew player %s successfully \n", player.Name)

	return playerCreated.ToPlayerResponse()
}

func (pc PlayerController) GetAll() []responses.PlayerResponse {
	return buildPlayerResponseAsArray(pc.playerService.GetAll())
}

func (pc PlayerController) GetById(id int16) responses.PlayerResponse {
	return pc.playerService.GetPlayerById(id).ToPlayerResponse()
}

func buildPlayerResponseAsArray(players []models.Player) []responses.PlayerResponse {
	playersResponses := make([]responses.PlayerResponse, 0)
	for _, player := range players {
		playersResponses = append(playersResponses, player.ToPlayerResponse())
	}
	return playersResponses
}
