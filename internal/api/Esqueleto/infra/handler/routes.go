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

	e.POST("/api/v1/besigabi/creacionTurnos", h.CrearCreacionTurnos)
	e.GET("/api/v1/besigabi/creacionTurnos", h.BuscarCreacionTurnos)
	// Otras rutas...
}
