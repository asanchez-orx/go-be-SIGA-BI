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

	e.POST("/api/v1/besigabi/creacionTurnos", h.CrearTurno)                                       //OK
	e.GET("/api/v1/besigabi/creacionTurnos/tipoDocumento", h.BuscarTipoDocumento)                 //OK
	e.POST("/api/v1/besigabi/creacionTurnos/companias", h.BuscarCompania)                         ///OK
	e.GET("/api/v1/besigabi/creacionTurnos/verificarConfigCompanias", h.VerificarConfigCompanias) //OK
	e.POST("/api/v1/besigabi/creacionTurnos/tipoServicio", h.BuscarTipoServicio)                  //OK
	e.POST("/api/v1/besigabi/creacionTurnos/modulos", h.BuscarModulo)                             //OK
	e.GET("/api/v1/besigabi/creacionTurnos/sedes", h.BuscarSedes)                                 //OK
	e.POST("/api/v1/besigabi/creacionTurnos/tipoTurno", h.BuscarTipoTurno)                        //OK
	e.GET("/api/v1/besigabi/creacionTurnos/confirmarConfigSedes", h.ConfirmarConfigSedes)         //ok
	e.GET("/api/v1/besigabi/creacionTurnos/cargarConfigLIS", h.CargarConfigLIS)                   //ok

}
