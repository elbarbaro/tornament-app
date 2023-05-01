package controllers

import (
	"fmt"

	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/services"
)

type TeamController struct {
	teamService interfaces.ITeamService
}

func NewTeamController() TeamController {
	return TeamController{
		teamService: services.NewTeamService(),
	}
}

func (tc TeamController) CreateTeam(name string, coach models.Coach) models.Team {
	fmt.Printf("Creating a team %s...\n", name)
	team := models.Team{
		Name:  name,
		Coach: coach,
	}
	teamCreated := tc.teamService.CreateTeam(team)
	return teamCreated
}

func (tc TeamController) GetAll() []models.Team {
	return tc.teamService.GetAll()
}

func (tc TeamController) GetTeamById(id int16) models.Team {
	return tc.teamService.GetTeamById(id)
}
