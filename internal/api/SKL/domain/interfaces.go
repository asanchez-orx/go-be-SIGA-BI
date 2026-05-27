package domain

import "context"

type SKLUseCase interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
	GetServiciosSiga(context.Context, ServiciosSigaRequest) (ServiciosSigaResponse, error)
	GetTurnosDisponibles(context.Context, TurnosDisponiblesRequest) (TurnosDisponiblesResponse, error)
	GetTurnosDisponiblesConOrden(context.Context, TurnosDisponiblesRequest) (TurnosDisponiblesResponse, error)
	GetSedesUsuario(context.Context, SedesUsuarioRequest) (SedesUsuarioResponse, error)
	ConsumirCredenciales(context.Context, ConsumirCredencialesRequest) (ConsumirCredencialesResponse, error)
}

type SKLRepository interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
	GetServiciosSiga(context.Context, ServiciosSigaRequest) (ServiciosSigaResponse, error)
	GetTurnosDisponibles(context.Context, TurnosDisponiblesRequest) (TurnosDisponiblesResponse, error)
	GetTurnosDisponiblesConOrden(context.Context, TurnosDisponiblesRequest) (TurnosDisponiblesResponse, error)
	GetSedesUsuario(context.Context, SedesUsuarioRequest) (SedesUsuarioResponse, error)
	ConsumirCredenciales(context.Context, ConsumirCredencialesRequest) (ConsumirCredencialesResponse, error)
}
