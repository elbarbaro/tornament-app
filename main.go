package main

import (
	"fmt"

	"elbarbaro.com/tornament-app/controllers"
)

func main() {
	fmt.Print("Hello World\n\n")

	coachController := controllers.NewCoachController()
	teamController := controllers.NewTeamController()
	playerController := controllers.NewPlayerController()

	malayo := coachController.CreateCoach("Benjamin Mora", 40)
	paunovich := coachController.CreateCoach("Velikoj Paunovich", 70)
	vucetich := coachController.CreateCoach("Victor Vucetich", 85)

	atlas := teamController.CreateTeam("Atlas", malayo)
	chivas := teamController.CreateTeam("Guadalajara", paunovich)
	rayados := teamController.CreateTeam("Monterrey", vucetich)

	fmt.Println(atlas)
	fmt.Println(chivas)
	fmt.Println(rayados)

	quinones := playerController.CreatePlayer("Julian Quiñones", 33, "CO", "DEL", "-")
	furch := playerController.CreatePlayer("Julio Furch", 9, "AR", "DEL", "-")
	camilo := playerController.CreatePlayer("Camilo Vargas", 12, "CO", "POR", "C")

	atlas.AddPlayer(quinones)
	atlas.AddPlayer(furch)
	atlas.AddPlayer(camilo)
	atlas.ShowPlayers()

	fmt.Println("\nGet all players...")
	players := playerController.GetAll()
	fmt.Println(players)
	fmt.Println("\nGet player where id is :")
	playerId1 := playerController.GetById(1)
	fmt.Println(playerId1)

	fmt.Println("\nGet all teams...")
	teams := teamController.GetAll()
	fmt.Println(teams)
	fmt.Println("\nGet team where id is 1:")
	teamId1 := teamController.GetTeamById(1)
	fmt.Println(teamId1)

	fmt.Println("\nGet all coachs...")
	coachs := coachController.GetAll()
	fmt.Println(coachs)

	fmt.Println("\nGet coach where id is 1:")
	coachId1 := coachController.GetCoachById(1)
	fmt.Println(coachId1)
}
