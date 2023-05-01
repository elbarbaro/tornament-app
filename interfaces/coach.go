package interfaces

import "elbarbaro.com/tornament-app/models"

type ICoachService interface {
	CreateCoach(coach models.Coach) models.Coach
	GetAll() []models.Coach
	GetCoachById(id int16) models.Coach
}

type ICoachRepository interface {
	Save(coach models.Coach) models.Coach
	FindAll() []models.Coach
	FindById(id int16) models.Coach
}
