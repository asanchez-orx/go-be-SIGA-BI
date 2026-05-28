package handler

import (
	"net/http"

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
