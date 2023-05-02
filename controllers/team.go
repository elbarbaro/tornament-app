package controllers

import (
	"fmt"

	"elbarbaro.com/tornament-app/handlers/requests"
	"elbarbaro.com/tornament-app/handlers/responses"
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

func (tc TeamController) CreateTeam(teamRequest requests.TeamRequest) models.Team {
	fmt.Printf("Creating a team %s...\n", teamRequest.Name)
	team := models.Team{
		Name: teamRequest.Name,
		Logo: teamRequest.Logo,
	}
	teamCreated := tc.teamService.CreateTeam(team)
	return teamCreated
}

func (tc TeamController) GetAll() []responses.TeamResponse {
	return buildTeamResponseAsArray(tc.teamService.GetAll())
}

func (tc TeamController) GetTeamById(id int16) models.Team {
	return tc.teamService.GetTeamById(id)
}

func buildTeamResponseAsArray(teams []models.Team) []responses.TeamResponse {
	teamsResponse := make([]responses.TeamResponse, 0)
	for _, team := range teams {
		teamsResponse = append(teamsResponse, team.ToTeamResponse())
	}
	return teamsResponse
}
