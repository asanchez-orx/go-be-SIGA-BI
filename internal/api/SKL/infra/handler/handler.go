package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"develop.private/CLTech/besigabi/internal/api/SKL/domain"
	"develop.private/CLTech/besigabi/libs/crypto"
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

// @Summary	Obtener servicios SIGA
// @Description	Consulta los servicios SIGA disponibles
// @Tags	SKL
// @Accept	json
// @Param	request	body	domain.ServiciosSigaRequest	true	"Parámetros de la consulta de servicios"
// @Produce	json
// @Success	200	{object}	domain.ServiciosSigaResponse
// @Success	204	{object}	nil "Sin servicios"
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/serviciosSiga [post]
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

// @Summary	Obtener turnos disponibles
// @Description	Consulta los turnos disponibles en LIS
// @Tags	SKL
// @Accept	json
// @Param	request	body	domain.TurnosDisponiblesRequest	true	"Parámetros de la consulta"
// @Produce	json
// @Success	200	{object}	domain.TurnosDisponiblesResponse
// @Success	204	{object}	nil "Sin turnos"
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/turnosDisponiblesLis [post]
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

// @Summary	Obtener sedes por usuario
// @Description	Consulta las sedes disponibles para un usuario
// @Tags	SKL
// @Accept	json
// @Param	request	body	domain.SedesUsuarioRequest	true	"Parámetros de usuario"
// @Produce	json
// @Success	200	{object}	domain.SedesUsuarioResponse
// @Success	204	{object}	nil "Sin sedes"
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/sedesUsuario [post]
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

// @Summary	Consumir credenciales
// @Description	Valida las credenciales de un usuario
// @Tags	SKL
// @Accept	json
// @Param	request	body	domain.ConsumirCredencialesRequest	true	"Credenciales"
// @Produce	json
// @Success	200	{object}	domain.ConsumirCredencialesResponse
// @Failure	400	{object}	echo.HTTPError
// @Failure	401	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/consumirCredenciales [post]
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

// @Summary	Obtener turnos disponibles con orden
// @Description	Consulta turnos disponibles asociados a una orden
// @Tags	SKL
// @Accept	json
// @Param	request	body	domain.TurnosDisponiblesRequest	true	"Parámetros de la consulta"
// @Produce	json
// @Success	200	{object}	domain.TurnosDisponiblesResponse
// @Success	204	{object}	nil "Sin turnos"
// @Failure	400	{object}	echo.HTTPError
// @Failure	500	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/turnosDisponibles [post]
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

// @Summary	Desencriptar payload dinámico
// @Description	Recibe un JSON dinámico, desencripta sus valores string y lo devuelve
// @Tags	SKL
// @Accept	json
// @Produce	json
// @Success	200	{object}	map[string]interface{}
// @Failure	400	{object}	echo.HTTPError
// @Router	/api/v1/besigabi/skl/desencriptNT [post]
func (h handler) DesencriptNT(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	encrypter := crypto.NewEncrypter("104F")
	response := make(map[string]interface{})

	for k, v := range payload {
		if strVal, ok := v.(string); ok {
			response[k] = encrypter.Decrypt(strVal)
		} else {
			response[k] = v // Keep original if not a string
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (h handler) EncriptNT(c echo.Context) error {
    var payload map[string]interface{}
    if err := c.Bind(&payload); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    encrypter := crypto.NewEncrypter("104F")
    response := make(map[string]interface{})
    for k, v := range payload {
        if strVal, ok := v.(string); ok {
            response[k] = encrypter.Encrypt(strVal)
        } else {
            response[k] = v // Keep original if not a string
        }
    }
    return c.JSON(http.StatusOK, response)
}
