package app

import (
	"context"

	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/domain"
)

type CreacionTurnosApp struct {
	repository domain.CreacionTurnosRepository
}

func NewCreacionTurnosApp(repository domain.CreacionTurnosRepository) *CreacionTurnosApp {
	return &CreacionTurnosApp{
		repository: repository,
	}
}

func (a *CreacionTurnosApp) BuscarCreacionTurnosService(ctx context.Context, req domain.CreacionTurnosRequest) (domain.CreacionTurnosesResponse, error) {
	return domain.CreacionTurnosesResponse{}, nil
}

func (a *CreacionTurnosApp) CrearCreacionTurnosService(ctx context.Context, req domain.CreacionTurnosRequest) error {
	return nil
}
