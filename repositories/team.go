package repositories

import (
	"elbarbaro.com/tornament-app/interfaces"
	"elbarbaro.com/tornament-app/models"
)

type teamRepository struct {
	db []models.Team
}

func NewTeamRepository() interfaces.ITeamRepository {
	return &teamRepository{}
}

func (tr *teamRepository) Save(team models.Team) models.Team {
	team.Id = int16(len(tr.db)) + 1
	tr.db = append(tr.db, team)
	return team
}

func (tr *teamRepository) FindAll() []models.Team {
	return tr.db
}

func (tr *teamRepository) FindById(id int16) models.Team {
	var teamFinded models.Team
	for _, team := range tr.db {
		if team.Id == id {
			teamFinded = team
			break
		}
	}
	return teamFinded
}
