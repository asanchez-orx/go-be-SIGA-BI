package handler

import (
	"develop.private/CLTech/besigabi/internal/api/SKL/app"
	"develop.private/CLTech/besigabi/internal/api/SKL/infra/mssql"
	"develop.private/CLTech/vulcano/infra/database"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := database.GetDatabase()
	repo := mssql.NewSKLRepo(db)
	appSvc := app.NewSKLApp(repo)
	h := newHandler(appSvc)

	e.POST("/api/v1/besigabi/skl/taquillas", h.GetTaquillas)
	e.POST("/api/v1/besigabi/skl/serviciosSiga", h.GetServiciosSiga)
	e.POST("/api/v1/besigabi/skl/turnosDisponiblesLis", h.GetTurnosDisponibles)
	e.POST("/api/v1/besigabi/skl/turnosDisponibles", h.GetTurnosDisponiblesConOrden)
	e.POST("/api/v1/besigabi/skl/sedesUsuario", h.GetSedesUsuario)
	e.POST("/api/v1/besigabi/skl/consumirCredenciales", h.ConsumirCredenciales)
}
