package controllers

import (
	"elbarbaro.com/tornament-app/handlers/requests"
	"elbarbaro.com/tornament-app/handlers/responses"
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/services"
)

type CoachController struct {
	coachService interfaces.ICoachService
}

func NewCoachController() CoachController {
	return CoachController{
		coachService: services.NewCoachService(),
	}
}

func (c CoachController) CreateCoach(coachRequest requests.CoachRequest) responses.CoachResponse {
	coach := models.Coach{
		Name:       coachRequest.Name,
		Percentage: coachRequest.Percentage,
	}
	coachCreated := c.coachService.CreateCoach(coach)

	return coachCreated.ToCoachResponse()
}

func (c CoachController) GetAll() []responses.CoachResponse {
	return buildCoachResponsesAsArray(c.coachService.GetAll())
}

func (c CoachController) GetCoachById(id int16) models.Coach {
	return c.coachService.GetCoachById(id)
}

func buildCoachResponsesAsArray(coachs []models.Coach) []responses.CoachResponse {
	coachsResponse := make([]responses.CoachResponse, 0)
	for _, coach := range coachs {
		coachsResponse = append(coachsResponse, coach.ToCoachResponse())
	}
	return coachsResponse
}
