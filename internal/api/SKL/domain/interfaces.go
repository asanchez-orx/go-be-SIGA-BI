package domain

import "context"

type SKLUseCase interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
	GetServiciosSiga(context.Context, ServiciosSigaRequest) (ServiciosSigaResponse, error)
}

type SKLRepository interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
	GetServiciosSiga(context.Context, ServiciosSigaRequest) (ServiciosSigaResponse, error)
}
