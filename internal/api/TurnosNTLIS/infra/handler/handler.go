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

func (h *handler) BuscarTaquillasXServicio(c echo.Context) error {

	sedeID, err := strconv.Atoi(c.Param("sede"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "sede inválida",
		})
	}

	servicioID, err := strconv.Atoi(c.Param("servicio"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "servicio inválido",
		})
	}

	response, err := h.turnosNTLISApp.BuscarTaquillasXServicio(
		c.Request().Context(),
		sedeID,
		servicioID,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BuscarMotivosTaquilla(c echo.Context) error {

	response, err := h.turnosNTLISApp.BuscarMotivosTaquilla(c.Request().Context())
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ActualizarEstadoTaquilla(c echo.Context) error {

	var req domain.TaquillasEstadoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "invalid request body",
		})
	}

	response, err := h.turnosNTLISApp.ActualizarEstadoTaquilla(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) ActualizarEstadoAtencion(c echo.Context) error {

	var req domain.AtencionEstadoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "invalid request body",
		})
	}

	response, err := h.turnosNTLISApp.ActualizarEstadoAtencion(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BuscarMotivosCancelacion(c echo.Context) error {

	response, err := h.turnosNTLISApp.BuscarMotivosCancelacion(c.Request().Context())
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) BuscarServiciosDisponiblesParaTransferenciaService(c echo.Context) error {

	idSede, err := strconv.Atoi(c.Param("idSede"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "idSede inválido",
		})
	}

	idServicio, err := strconv.Atoi(c.Param("idServicio"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "idServicio inválido",
		})
	}

	idTurno, err := strconv.Atoi(c.Param("idTurno"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "idTurno inválido",
		})
	}

	response, err := h.turnosNTLISApp.BuscarServiciosDisponiblesParaTransferenciaService(c.Request().Context(), idSede, idServicio, idTurno)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) LlamadoTurno(c echo.Context) error {
	turno, err := strconv.Atoi(c.Param("turno"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": 400, "error": "turno inválido"})
	}
	servicioID, err := strconv.Atoi(c.Param("servicio"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": 400, "error": "servicio inválido"})
	}

	response, err := h.turnosNTLISApp.LlamadoTurnoService(c.Request().Context(), turno, servicioID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": 500, "error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) LlamadoTurnoPost(c echo.Context) error {
	var req domain.LlamadoTurnoPostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": 400, "error": "invalid request body"})
	}

	response, err := h.turnosNTLISApp.LlamadoTurnoPostService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": 500, "error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
