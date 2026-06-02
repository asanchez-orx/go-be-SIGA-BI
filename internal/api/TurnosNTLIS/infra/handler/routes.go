package handler

import (
	"develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/app"
	"develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/infra/mssql"
	"develop.private/CLTech/vulcano/infra/database"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := database.GetDatabase()
	repo := mssql.NewTurnosNTLISRepo(db)
	happ := app.NewTurnosNTLISApp(repo)
	h := newHandler(happ)

	e.GET("/api/branch/getAll", h.BuscarSedesNTService)
	e.GET("/api/service/getByBranch/:idSede", h.BuscarServiciosNTXSedeService)
	e.GET("/api/pointsOfCare", h.BuscarTaquillasNTService)
	e.GET("/api/taquillasxSedeServicio/:sede/:servicio", h.BuscarTaquillasXServicio)
	e.GET("/api/reasons/break", h.BuscarMotivosTaquilla)
	e.POST("/api/log", h.ActualizarEstadoTaquilla)
	e.POST("/api/logAtencion", h.ActualizarEstadoAtencion)
	e.GET("/api/reasons/cancel", h.BuscarMotivosCancelacion)
	e.GET("/api/reasons/transfer/:idSede/:idServicio/:idTurno", h.BuscarServiciosDisponiblesParaTransferenciaService)
	e.GET("/api/turns/call/:turno/:servicio", h.LlamadoTurno)
	e.POST("/api/turns/call", h.LlamadoTurnoPost)
	e.GET("/api/turns/daily/:sede/:servicio/:apellido/:nombre/:userName/:taquilla", h.BuscarTurnosDisponibles)
	e.POST("/api/transfers", h.TransferirTurno)
}
