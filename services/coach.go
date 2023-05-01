package services

import (
	"fmt"

	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
	"elbarbaro.com/tornament-app/repositories"
)

type coachService struct {
	coachRepository interfaces.ICoachRepository
}

func NewCoachService() interfaces.ICoachService {
	return coachService{
		coachRepository: repositories.NewCoachRepository(),
	}
}

func (cs coachService) CreateCoach(coach models.Coach) models.Coach {
	fmt.Printf("\nCreate new coach %s...\n", coach.Name)
	coachCreated := cs.coachRepository.Save(coach)
	fmt.Printf("\nCoach successfully create %s...\n", coach.Name)
	return coachCreated
}

func (cs coachService) GetAll() []models.Coach {
	return cs.coachRepository.FindAll()
}

func (cs coachService) GetCoachById(id int16) models.Coach {
	return cs.coachRepository.FindById(id)
}
