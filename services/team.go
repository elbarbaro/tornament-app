package services

import (
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/repositories"
)

type teamService struct {
	teamRepository interfaces.ITeamRepository
}

func NewTeamService() interfaces.ITeamService {
	return &teamService{
		teamRepository: repositories.NewTeamRepository(),
	}
}

func (ts *teamService) CreateTeam(team models.Team) models.Team {
	return ts.teamRepository.Save(team)
}

func (ts *teamService) GetAll() []models.Team {
	return ts.teamRepository.FindAll()
}

func (ts *teamService) GetTeamById(id int16) models.Team {
	return ts.teamRepository.FindById(id)
}
