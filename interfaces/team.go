package interfaces

import "elbarbaro.com/tornament-app/models"

type ITeamService interface {
	CreateTeam(models.Team) models.Team
	GetAll() []models.Team
	GetTeamById(id int16) models.Team
}

type ITeamRepository interface {
	Save(models.Team) models.Team
	FindAll() []models.Team
	FindById(id int16) models.Team
}
