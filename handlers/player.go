package handlers

import (
	"net/http"
	"strconv"

	"elbarbaro.com/tornament-app/controllers"
	"elbarbaro.com/tornament-app/handlers/requests"
	"github.com/labstack/echo/v4"
)

type playerHandler struct {
	playerController controllers.PlayerController
}

func NewPlayerHandler() *playerHandler {
	return &playerHandler{
		playerController: controllers.NewPlayerController(),
	}
}

func (ph *playerHandler) CreatePlayer(c echo.Context) error {
	var playerRequest requests.PlayerRequest

	if err := c.Bind(&playerRequest); err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}

	playerResponse := ph.playerController.CreatePlayer(playerRequest)
	return c.JSON(http.StatusCreated, playerResponse)
}

func (ph *playerHandler) GetAll(c echo.Context) error {
	playersResponse := ph.playerController.GetAll()
	return c.JSON(http.StatusOK, playersResponse)
}

func (ph *playerHandler) GetPlayerById(c echo.Context) error {
	var playerId int64
	var err error

	if playerId, err = strconv.ParseInt(c.Param("id"), 10, 16); err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}

	playerResponse := ph.playerController.GetById(int16(playerId))
	return c.JSON(http.StatusOK, playerResponse)
}
