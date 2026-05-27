package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"develop.private/CLTech/besigabi/internal/api/SKL/domain"
)

type handler struct {
	sklApp domain.SKLUseCase
}

func newHandler(app domain.SKLUseCase) *handler {
	return &handler{
		sklApp: app,
	}
}

// @Summary	Obtener taquillas por sede y módulo
// @Description	Consulta las taquillas activas para una sede y módulo específico
// @Tags	Taquillas
// @Accept	json
// @Param	request	body	domain.TaquillasRequest	true	"Parámetros de la consulta de taquillas"
// @Produce	json
// @Success	200	{object}	domain.TaquillasResponse
// @Success	204	{object}	nil "Sin taquillas"
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/taquillas [post]
func (h handler) GetTaquillas(c echo.Context) error {
	req := domain.TaquillasRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.GetTaquillas(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetServiciosSiga(c echo.Context) error {
	req := domain.ServiciosSigaRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.GetServiciosSiga(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetTurnosDisponibles(c echo.Context) error {
	req := domain.TurnosDisponiblesRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.GetTurnosDisponibles(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetSedesUsuario(c echo.Context) error {
	req := domain.SedesUsuarioRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.GetSedesUsuario(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) ConsumirCredenciales(c echo.Context) error {
	req := domain.ConsumirCredencialesRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.ConsumirCredenciales(c.Request().Context(), req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Credenciales incorrectas o usuario no encontrado")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetTurnosDisponiblesConOrden(c echo.Context) error {
	req := domain.TurnosDisponiblesRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.sklApp.GetTurnosDisponiblesConOrden(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}
