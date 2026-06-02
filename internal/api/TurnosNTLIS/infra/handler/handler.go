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

// @Summary		Obtener todas las sedes
// @Description	Obtiene la lista de todas las sedes estructuradas
// @Tags		TurnosNTLIS
// @Produce		json
// @Success		200	{object}	domain.SedesNTResponse
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/branch/getAll [get]
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

// @Summary		Obtener servicios por sede
// @Description	Obtiene los servicios disponibles para una sede específica
// @Tags		TurnosNTLIS
// @Produce		json
// @Param		idSede path int true "ID de la Sede"
// @Success		200	{object}	domain.ServiciosNTXSedeResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/service/getByBranch/{idSede} [get]
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

// @Summary		Obtener puntos de atención
// @Description	Obtiene todos los puntos de atención disponibles
// @Tags		TurnosNTLIS
// @Produce		json
// @Success		200	{object}	domain.TaquillaNTResponse
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/pointsOfCare [get]
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

// @Summary		Obtener taquillas por sede y servicio
// @Description	Obtiene las taquillas filtradas por sede y servicio
// @Tags		TurnosNTLIS
// @Produce		json
// @Param		sede path int true "ID de la Sede"
// @Param		servicio path int true "ID del Servicio"
// @Success		200	{object}	domain.TaquillaxSedexServicioResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/taquillasxSedeServicio/{sede}/{servicio} [get]
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

// @Summary		Obtener motivos de descanso
// @Description	Obtiene la lista de motivos de descanso (break) para las taquillas
// @Tags		TurnosNTLIS
// @Produce		json
// @Success		200	{object}	domain.TaquillaMotivoDescansoResponse
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/reasons/break [get]
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

// @Summary		Actualizar estado de la taquilla
// @Description	Actualiza el estado log de una taquilla
// @Tags		TurnosNTLIS
// @Accept		json
// @Produce		json
// @Param		request body domain.TaquillasEstadoRequest true "Datos del estado de taquilla"
// @Success		200	{object}	domain.TaquillasEstadoResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/log [post]
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

// @Summary		Actualizar estado de atención
// @Description	Actualiza el estado de la atención de un turno
// @Tags		TurnosNTLIS
// @Accept		json
// @Produce		json
// @Param		request body domain.AtencionEstadoRequest true "Datos del estado de atención"
// @Success		200	{object}	domain.AtencionEstadoResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/logAtencion [post]
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

// @Summary		Obtener motivos de cancelación
// @Description	Obtiene la lista de motivos de cancelación
// @Tags		TurnosNTLIS
// @Produce		json
// @Success		200	{object}	domain.Response
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/reasons/cancel [get]
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

// @Summary		Obtener servicios para transferencia
// @Description	Obtiene servicios disponibles para la transferencia de un turno
// @Tags		TurnosNTLIS
// @Produce		json
// @Param		idSede path int true "ID de la Sede"
// @Param		idServicio path int true "ID del Servicio"
// @Param		idTurno path int true "ID del Turno"
// @Success		200	{object}	domain.ResponseTransfer
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/reasons/transfer/{idSede}/{idServicio}/{idTurno} [get]
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

// @Summary		Verificar llamado de turno
// @Description	Verifica si un turno ha sido llamado
// @Tags		TurnosNTLIS
// @Produce		json
// @Param		turno path int true "Número de Turno"
// @Param		servicio path int true "ID de Servicio"
// @Success		200	{object}	domain.LlamadoTurnoResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/turns/call/{turno}/{servicio} [get]
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

// @Summary		Llamar un turno
// @Description	Actualiza el estado de un turno a llamado
// @Tags		TurnosNTLIS
// @Accept		json
// @Produce		json
// @Param		request body domain.LlamadoTurnoPostRequest true "Datos para llamar turno"
// @Success		200	{object}	domain.LlamadoTurnoPostResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/turns/call [post]
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

// @Summary		Obtener turnos diarios disponibles
// @Description	Obtiene todos los turnos diarios disponibles para un servicio de una sede
// @Tags		TurnosNTLIS
// @Produce		json
// @Param		sede path int true "ID de Sede"
// @Param		servicio path int true "ID de Servicio"
// @Param		apellido path string false "Apellido del paciente"
// @Param		nombre path string false "Nombre del paciente"
// @Param		userName path string false "Nombre de usuario"
// @Param		taquilla path string false "ID de la taquilla"
// @Success		200	{object}	domain.TurnoConsultaResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/turns/daily/{sede}/{servicio}/{apellido}/{nombre}/{userName}/{taquilla} [get]
func (h *handler) BuscarTurnosDisponibles(c echo.Context) error {
	sedeID, err := strconv.Atoi(c.Param("sede"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": 400, "error": "sede inválida"})
	}

	servicioID, err := strconv.Atoi(c.Param("servicio"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": 400, "error": "servicio inválido"})
	}

	response, err := h.turnosNTLISApp.BuscarTurnosDisponiblesService(c.Request().Context(), sedeID, servicioID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": 500, "error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary		Transferir turno a otro servicio
// @Description	Transfiere un turno creando uno nuevo asociado al servicio destino y actualizando el viejo
// @Tags		TurnosNTLIS
// @Accept		json
// @Produce		json
// @Param		request body domain.TransferRequest true "Datos de la transferencia"
// @Success		200	{object}	domain.TransferResponse
// @Failure		400	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router		/api/transfers [post]
func (h *handler) TransferirTurno(c echo.Context) error {
	var req domain.TransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"error":  "invalid request body",
		})
	}

	response, err := h.turnosNTLISApp.TransferirTurnoService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
