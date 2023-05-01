package controllers

import (
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/services"
)

type coachController struct {
	coachService interfaces.ICoachService
}

func NewCoachController() coachController {
	return coachController{
		coachService: services.NewCoachService(),
	}
}

func (c coachController) CreateCoach(name string, percentage float32) models.Coach {
	coach := models.Coach{
		Name:       name,
		Percentage: percentage,
	}
	coachCreated := c.coachService.CreateCoach(coach)

	return coachCreated
}

func (c coachController) GetAll() []models.Coach {
	return c.coachService.GetAll()
}

func (c coachController) GetCoachById(id int16) models.Coach {
	return c.coachService.GetCoachById(id)
}
