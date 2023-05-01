package interfaces

import "elbarbaro.com/tornament-app/models"

type IPlayeService interface {
	CreatePlayer(player models.Player) models.Player
	GetAll() []models.Player
	GetPlayerById(id int16) models.Player
}

type IPlayerRepository interface {
	Save(player models.Player) models.Player
	FindAll() []models.Player
	FindById(id int16) models.Player
}
