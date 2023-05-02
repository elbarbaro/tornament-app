package handlers

import (
	"net/http"
	"strconv"

	"elbarbaro.com/tornament-app/controllers"
	"elbarbaro.com/tornament-app/handlers/requests"
	"github.com/labstack/echo/v4"
)

type coachHandler struct {
	coachController controllers.CoachController
}

func NewCoachHandler() *coachHandler {
	return &coachHandler{
		coachController: controllers.NewCoachController(),
	}
}

func (ch *coachHandler) CreateCoach(c echo.Context) error {
	var coachRequest requests.CoachRequest

	if err := c.Bind(&coachRequest); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	coachResponse := ch.coachController.CreateCoach(coachRequest)

	return c.JSON(http.StatusCreated, coachResponse)
}

func (ch *coachHandler) GetAll(c echo.Context) error {
	response := ch.coachController.GetAll()
	return c.JSON(http.StatusOK, response)
}

func (ch *coachHandler) GetCoachById(c echo.Context) error {
	var coachId int64
	var err error

	if coachId, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	coachResponse := ch.coachController.GetCoachById(int16(coachId))

	return c.JSON(http.StatusOK, coachResponse)
}
