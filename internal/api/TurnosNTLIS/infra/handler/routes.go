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
}
