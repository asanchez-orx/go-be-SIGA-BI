package app

import (
	"context"

	"develop.private/CLTech/besigabi/internal/api/SKL/domain"
)

type SKLApp struct {
	repository domain.SKLRepository
}

func NewSKLApp(repository domain.SKLRepository) *SKLApp {
	return &SKLApp{
		repository: repository,
	}
}

func (a *SKLApp) GetTaquillas(ctx context.Context, req domain.TaquillasRequest) (domain.TaquillasResponse, error) {
	return a.repository.GetTaquillas(ctx, req)
}
