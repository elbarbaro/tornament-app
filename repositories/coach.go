package repositories

import (
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
)

type coachRepository struct {
	db []models.Coach
}

func NewCoachRepository() interfaces.ICoachRepository {
	return &coachRepository{}
}

func (cr *coachRepository) Save(coach models.Coach) models.Coach {
	coach.Id = int16(len(cr.db)) + 1
	cr.db = append(cr.db, coach)
	return coach
}

func (cr *coachRepository) FindAll() []models.Coach {
	return cr.db
}

func (cr *coachRepository) FindById(id int16) models.Coach {
	var coachFinded models.Coach
	for _, coach := range cr.db {
		if coach.Id == id {
			coachFinded = coach
			break
		}
	}
	return coachFinded
}
