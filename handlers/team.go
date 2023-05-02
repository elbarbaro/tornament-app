package handlers

import (
	"net/http"
	"strconv"

	"elbarbaro.com/tornament-app/controllers"
	"elbarbaro.com/tornament-app/handlers/requests"
	"github.com/labstack/echo/v4"
)

type teamHandler struct {
	teamController controllers.TeamController
}

func NewTeamHandler() *teamHandler {
	return &teamHandler{
		teamController: controllers.NewTeamController(),
	}
}

func (th *teamHandler) CreateTeam(c echo.Context) error {

	var request requests.TeamRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdTeam := th.teamController.CreateTeam(request)

	return c.JSON(http.StatusCreated, createdTeam.ToTeamResponse())
}

func (th *teamHandler) GetAll(c echo.Context) error {
	teams := th.teamController.GetAll()
	return c.JSON(http.StatusOK, teams)
}

func (th *teamHandler) GetTeamById(c echo.Context) error {

	var teamId int64
	var err error

	if teamId, err = strconv.ParseInt(c.Param("id"), 10, 16); err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}

	team := th.teamController.GetTeamById(int16(teamId))

	return c.JSON(http.StatusOK, team.ToTeamResponse())
}
