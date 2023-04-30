package main

import (
	"fmt"

	"elbarbaro.com/tornament-app/models"
)

func main() {
	fmt.Println("Hello World")

	malayo := models.Coach{
		Id:   1,
		Name: "Benjamin Mora",
	}

	paunovich := models.Coach{
		Id:   2,
		Name: "Velikoj Paunovich",
	}

	vucetich := models.Coach{
		Id:   3,
		Name: "Victor Vucetich",
	}

	atlas := models.Team{
		Id:    1,
		Name:  "Atlas",
		Coach: malayo,
	}

	chivas := models.Team{
		Id:    2,
		Name:  "Guadalajara",
		Coach: paunovich,
	}

	rayados := models.Team{
		Id:    3,
		Name:  "Monterrey",
		Coach: vucetich,
	}

	fmt.Println(atlas)
	fmt.Println(chivas)
	fmt.Println(rayados)

	quinones := models.Player{
		Id:       1,
		Name:     "Julian Quiñones",
		Number:   33,
		Country:  "CO",
		Position: "DEL",
	}

	atlas.AddPlayer(quinones)

	atlas.ShowPlayers()

	clausura2023 := models.Tournament{
		Id:   1,
		Name: "Clausura 2023",
	}

	game1 := models.Game{
		Id:      1,
		TeamA:   atlas,
		TeamB:   chivas,
		Stadium: "Estadio Jalisco",
		Date:    "2023-10-01",
	}

	game2 := models.Game{
		Id:      2,
		TeamA:   rayados,
		TeamB:   atlas,
		Stadium: "Estadio BBVA",
		Date:    "2023-10-07",
	}

	game3 := models.Game{
		Id:      2,
		TeamA:   chivas,
		TeamB:   rayados,
		Stadium: "Estadio Akron",
		Date:    "2023-10-14",
	}

	clausura2023.AddGame(game1)
	clausura2023.AddGame(game2)
	clausura2023.AddGame(game3)

	clausura2023.ShowGames()

	resultGame1 := models.Score{
		Id:     1,
		Game:   game1,
		GoalsA: 1,
		GoalsB: 0,
	}

	resultGame2 := models.Score{
		Id:     2,
		Game:   game2,
		GoalsA: 2,
		GoalsB: 3,
	}

	resultGame3 := models.Score{
		Id:     3,
		Game:   game3,
		GoalsA: 1,
		GoalsB: 1,
	}

	fmt.Println(resultGame1)
	fmt.Println(resultGame2)
	fmt.Println(resultGame3)

	atlasTournament := models.TournamentTeam{
		Id:           1,
		Team:         atlas,
		Points:       21,
		WonMatches:   4,
		LostMatches:  3,
		DrawnMatches: 8,
	}

	chivasTournament := models.TournamentTeam{
		Id:           2,
		Team:         chivas,
		Points:       31,
		WonMatches:   6,
		LostMatches:  4,
		DrawnMatches: 4,
	}

	rayadosTournament := models.TournamentTeam{
		Id:           3,
		Team:         rayados,
		Points:       37,
		WonMatches:   12,
		LostMatches:  2,
		DrawnMatches: 1,
	}

	clausura2023Leaderboard := models.Leaderboard{
		Tournament: clausura2023,
	}

	clausura2023Leaderboard.AddTournamentTeam(atlasTournament)
	clausura2023Leaderboard.AddTournamentTeam(chivasTournament)
	clausura2023Leaderboard.AddTournamentTeam(rayadosTournament)

	clausura2023Leaderboard.Show()
}
