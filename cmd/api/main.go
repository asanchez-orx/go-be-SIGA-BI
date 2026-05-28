package main

import (
	"fmt"
	"log/slog"
	"os"

	"develop.private/CLTech/vulcano/config"
	"develop.private/CLTech/vulcano/infra/database"
	vecho "develop.private/CLTech/vulcano/infra/echo"
	"develop.private/CLTech/vulcano/logger"
	"develop.private/CLTech/vulcano/service"

	sklHandler "develop.private/CLTech/besigabi/internal/api/SKL/infra/handler"
	TurnosNTLISHandler "develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/infra/handler"
	creacionTurnosHandler "develop.private/CLTech/besigabi/internal/api/creacionTurnos/infra/handler"
)

var Version = "Devel"
var BuildTime = "0000-00-00 00:00:00"
var CommitHash = "00000001"

// @title		API Backend de la aplicación besigabi
// @version		0.1.0
// @description	Descripción de este servidor backend
// @BasePath		/api/v1/
func main() {
	var cfgFile string
	if len(os.Args) > 1 {
		cfgFile = os.Args[1]
	} else {
		exit("No se ha especificado el archivo de configuración")
	}

	// Cargando configuración
	cfg, err := config.Read(cfgFile)
	if err != nil {
		exit(err.Error())
	}

	// Inicializar logger
	logger.Init(config.SlogLevel(cfg.Server.LogLevel), "besigabi", cfg.Server.LogDestination)
	slog.SetDefault(logger.Log)

	// Crear instancia de Echo
	e := vecho.NewEchoInstance(
		logger.Log,
		Version,
		BuildTime,
		CommitHash,
	)

	err = database.New(cfg.Database)
	if err != nil {
		exit(err.Error())
	}
	defer database.GetDatabase().Close()
	logger.Log.Info(
		fmt.Sprintf("Conexión a la base de datos %s:%d establecida correctamente", cfg.Database.Host, cfg.Database.Port),
	)

	creacionTurnosHandler.Routes(e)
	sklHandler.Routes(e)
	TurnosNTLISHandler.Routes(e)

	srv := vecho.EchoServer{
		App:  e,
		Addr: fmt.Sprintf(":%d", cfg.Server.Port),
		Log:  logger.Log,
	}

	err = service.RunAsService("besigabi", logger.Log, func() {
		service.RunGracefully(logger.Log, srv)
	})
	if err != nil {
		logger.Log.Error("Error al ejecutar servidor", "error", err)
	}
}

func exit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
