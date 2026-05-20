package domain

import "context"

type SKLUseCase interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
}

type SKLRepository interface {
	GetTaquillas(context.Context, TaquillasRequest) (TaquillasResponse, error)
}
