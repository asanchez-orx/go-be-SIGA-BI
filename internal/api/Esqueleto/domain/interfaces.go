package domain

import "context"

type CreacionTurnosUseCase interface {
	BuscarCreacionTurnosService(context.Context, CreacionTurnosRequest) (CreacionTurnosesResponse, error)
	CrearCreacionTurnosService(context.Context, CreacionTurnosRequest) error
}

type CreacionTurnosRepository interface {
	BuscarCreacionTurnos(context.Context, CreacionTurnosRequest) (CreacionTurnosesResponse, error)
	CrearCreacionTurnos(context.Context, CreacionTurnosRequest) error
}
