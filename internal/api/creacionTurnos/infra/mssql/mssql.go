package mssql

import (
	"context"

	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/domain"
	"develop.private/CLTech/vulcano/infra/database"
)

type CreacionTurnosRepo struct {
	db database.Database
}

func NewCreacionTurnosRepo(db database.Database) *CreacionTurnosRepo {
	return &CreacionTurnosRepo{
		db: db,
	}
}

func (r *CreacionTurnosRepo) BuscarCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) (domain.CreacionTurnosesResponse, error) {
	return domain.CreacionTurnosesResponse{}, nil
}

func (r *CreacionTurnosRepo) CrearCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) error {
	return nil
}
