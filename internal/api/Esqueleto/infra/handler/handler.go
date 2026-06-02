package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/domain"
)

type handler struct {
	creacionTurnosApp domain.CreacionTurnosUseCase
}

func newHandler(app domain.CreacionTurnosUseCase) *handler {
	return &handler{
		creacionTurnosApp: app,
	}
}

// @Summary		Resumen de lo que hace  el endpoint
// @Description	Descripción detallada del endpoint
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.CreacionTurnosRequest	true	"Parámetros de la consulta de las ordenes iniciales"
// @Produce			json
// @Success		200	{object}	domain.CreacionTurnosResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		404	{object}	map[string]interface{}
// @Router			/api/v1/besigabi/creacionTurnos [get]
func (h handler) BuscarCreacionTurnos(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// @Summary		Resumen de lo que hace  el endpoint
// @Description	Descripción detallada del endpoint
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.CreacionTurnosRequest	true	"Parámetros a crear en la BD"
// @Produce			json
// @Success		204	{object}	nil "No Content"
// @Failure		400	{object}	map[string]interface{}
// @Failure		404	{object}	map[string]interface{}
// @Router			/api/v1/besigabi/creacionTurnos [post]
func (h handler) CrearCreacionTurnos(c echo.Context) error {
	req := domain.CreacionTurnosRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
