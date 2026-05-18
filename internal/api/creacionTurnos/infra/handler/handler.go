package handler

import (
	"errors"
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
// @Failure		400	{object}	middleware.ClientError
// @Failure		404	{object}	middleware.ClientError
// @Router			/api/v1/besigabi/creacionTurnos [get]
func (h handler) BuscarCreacionTurnos(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// @Summary		Crea un turno
// @Description	Crea un turno en el sistema validando cupos y límites por módulo y origen de forma atómica
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.CrearTurnoRequest	true	"Parámetros para la creación del turno"
// @Produce			json
// @Success		201	{object}	domain.CrearTurnoResponse
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		429	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos [post]
func (h handler) CrearTurno(c echo.Context) error {
	req := domain.CrearTurnoRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.creacionTurnosApp.CrearTurnoService(c.Request().Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTurnLimitExceeded):
			return echo.NewHTTPError(http.StatusTooManyRequests, err.Error())
		case errors.Is(err, domain.ErrSedeNotFound),
			errors.Is(err, domain.ErrCompaniaNotFound),
			errors.Is(err, domain.ErrTipoTurnoNotFound),
			errors.Is(err, domain.ErrServicioNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, res)
}

// @Summary		Obtiene el listado de tipos de documento
// @Description	Consulta la tabla ENT0024 para obtener los tipos de documento activos (C05 = 1)
// @Tags			CreacionTurnos
// @Produce			json
// @Success		200	{object}	domain.TipoDocumentosesResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		500	{object}	error
// @Router			/api/v1/besigabi/creacionTurnos/tipoDocumento [get]
func (h handler) BuscarTipoDocumento(c echo.Context) error {
	res, err := h.creacionTurnosApp.BuscarTipoDocumentoService(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Obtiene el listado de compañías
// @Description	Consulta las tablas ENT5802 y ENT5814 para obtener las compañías asociadas a una sede
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.CompaniaRequest	true	"Parámetros de la consulta de compañías"
// @Produce			json
// @Success		200	{object}	domain.CompaniasResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/companias [post]
func (h handler) BuscarCompania(c echo.Context) error {
	req := domain.CompaniaRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.creacionTurnosApp.BuscarCompaniaService(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Verifica la configuración de compañías
// @Description	Consulta la tabla ENT5803 para verificar si se manejan empresas
// @Tags			CreacionTurnos
// @Produce			json
// @Success		200	{object}	domain.ConfigCompaniasResponse
// @Success		204	{object}	nil "Sin Configuración"
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/verificarConfigCompanias [get]
func (h handler) VerificarConfigCompanias(c echo.Context) error {
	res, err := h.creacionTurnosApp.VerificarConfigCompaniasService(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if res.ManejaEmpresas == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Obtiene el listado de tipos de servicio
// @Description	Consulta las tablas ENT5800 y ENT5815 para obtener los tipos de servicio
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.TipoServicioRequest	true	"Parámetros de la consulta de tipos de servicio"
// @Produce			json
// @Success		200	{object}	domain.TipoServiciosResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/tipoServicio [post]
func (h handler) BuscarTipoServicio(c echo.Context) error {
	req := domain.TipoServicioRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.creacionTurnosApp.BuscarTipoServicioService(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Obtiene el listado de módulos
// @Description	Consulta la tabla ENT5818 para obtener los módulos por sede
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.ModuloRequest	true	"Parámetros de la consulta de módulos"
// @Produce			json
// @Success		200	{object}	domain.ModulosResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/modulos [post]
func (h handler) BuscarModulo(c echo.Context) error {
	req := domain.ModuloRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.creacionTurnosApp.BuscarModuloService(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Obtiene el listado de sedes
// @Description	Consulta la tabla ENT0021 para obtener las sedes
// @Tags			CreacionTurnos
// @Produce			json
// @Success		200	{object}	domain.SedesResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/sedes [get]
func (h handler) BuscarSedes(c echo.Context) error {
	res, err := h.creacionTurnosApp.BuscarSedesService(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Obtiene el listado de tipos de turno
// @Description	Consulta las tablas ENT5816 y ENT5810 para obtener los tipos de turno por servicio
// @Tags			CreacionTurnos
// @Accept			json
// @Param			request	body	domain.TipoTurnoRequest	true	"Parámetros de la consulta de tipos de turno"
// @Produce			json
// @Success		200	{object}	domain.TipoTurnosResponse
// @Success		204	{object}	nil "Sin Data"
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/tipoTurno [post]
func (h handler) BuscarTipoTurno(c echo.Context) error {
	req := domain.TipoTurnoRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.creacionTurnosApp.BuscarTipoTurnoService(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(res) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary		Confirmar configuración de sedes
// @Description	Verifica si se manejan sedes y si no, retorna la sede por defecto
// @Tags			CreacionTurnos
// @Produce			json
// @Success		200	{object}	interface{}
// @Success		204	{object}	nil "No se encontró sede por Default"
// @Failure		500	{object}	echo.HTTPError
// @Router			/api/v1/besigabi/creacionTurnos/confirmarConfigSedes [get]
func (h handler) ConfirmarConfigSedes(c echo.Context) error {
	res, err := h.creacionTurnosApp.ConfirmarConfigSedesService(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en la consulta")
	}

	if res == nil {
		return c.JSON(http.StatusNoContent, "No se encontro sede por Default")
	}

	return c.JSON(http.StatusOK, res)
}
