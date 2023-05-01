package repositories

import (
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
)

type playerRepository struct {
	db []models.Player
}

func NewPlayerRepository() interfaces.IPlayerRepository {
	return &playerRepository{}
}

func (pr *playerRepository) Save(player models.Player) models.Player {
	player.Id = int16(len(pr.db)) + 1
	pr.db = append(pr.db, player)
	return player
}

func (pr *playerRepository) FindAll() []models.Player {
	return pr.db
}

func (pr *playerRepository) FindById(id int16) models.Player {
	var playerFinded models.Player
	for _, player := range pr.db {
		if player.Id == id {
			playerFinded = player
			break
		}
	}
	return playerFinded
}
