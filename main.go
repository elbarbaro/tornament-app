package main

import (
	"elbarbaro.com/tornament-app/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	echo := echo.New()

	teamHandler := handlers.NewTeamHandler()
	playerHandler := handlers.NewPlayerHandler()
	coachHandler := handlers.NewCoachHandler()

	echo.GET("/teams", teamHandler.GetAll)
	echo.POST("/teams", teamHandler.CreateTeam)
	echo.GET("/teams/:id", teamHandler.GetTeamById)

	echo.GET("/players", playerHandler.GetAll)
	echo.POST("/players", playerHandler.CreatePlayer)
	echo.GET("/players/:id", playerHandler.GetPlayerById)

	echo.GET("/coachs", coachHandler.GetAll)
	echo.POST("/coachs", coachHandler.CreateCoach)
	echo.GET("/coachs/:id", coachHandler.GetCoachById)

	echo.Logger.Fatal(echo.Start(":1323"))
}
