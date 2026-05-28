package handler

import (
	"net/http"
	"strconv"

	"develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/domain"

	"github.com/labstack/echo/v4"
)

type handler struct {
	turnosNTLISApp domain.TurnosNTLISUseCase
}

func newHandler(app domain.TurnosNTLISUseCase) *handler {
	return &handler{
		turnosNTLISApp: app,
	}
}

func (h *handler) BuscarSedesNTService(c echo.Context) error {

	response, err := h.turnosNTLISApp.BuscarSedesNTService(c.Request().Context())
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BuscarServiciosNTXSedeService(c echo.Context) error {

	sedeID, err := strconv.Atoi(c.Param("idSede"))
	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "idSede inválido",
		})
	}

	response, err := h.turnosNTLISApp.BuscarServiciosNTXSedeService(
		c.Request().Context(),
		sedeID,
	)

	if err != nil {

		println(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BuscarTaquillasNTService(c echo.Context) error {

	response, err := h.turnosNTLISApp.BuscarTaquillasNTService(c.Request().Context())
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
