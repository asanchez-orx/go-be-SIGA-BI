package handler

import (
	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/app"
	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/infra/mssql"
	"develop.private/CLTech/vulcano/infra/database"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := database.GetDatabase()
	repo := mssql.NewCreacionTurnosRepo(db)
	happ := app.NewCreacionTurnosApp(repo)
	h := newHandler(happ)

	e.POST("/api/v1/besigabi/creacionTurnos", h.CrearTurno)
	e.GET("/api/v1/besigabi/creacionTurnos", h.BuscarCreacionTurnos)
	e.GET("/api/v1/besigabi/creacionTurnos/tipoDocumento", h.BuscarTipoDocumento)
	e.POST("/api/v1/besigabi/creacionTurnos/companias", h.BuscarCompania)
	e.GET("/api/v1/besigabi/creacionTurnos/verificarConfigCompanias", h.VerificarConfigCompanias)
	e.POST("/api/v1/besigabi/creacionTurnos/tipoServicio", h.BuscarTipoServicio)
	e.POST("/api/v1/besigabi/creacionTurnos/modulos", h.BuscarModulo)
	e.GET("/api/v1/besigabi/creacionTurnos/sedes", h.BuscarSedes)
	e.POST("/api/v1/besigabi/creacionTurnos/tipoTurno", h.BuscarTipoTurno)
	e.GET("/api/v1/besigabi/creacionTurnos/confirmarConfigSedes", h.ConfirmarConfigSedes)
	e.GET("/api/v1/besigabi/creacionTurnos/cargarConfigLIS", h.CargarConfigLIS)
}
