package app

import (
	"context"

	"develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/domain"
)

type turnosNTLISApp struct {
	repository domain.TurnosNTLISRepository
}

func NewTurnosNTLISApp(
	repository domain.TurnosNTLISRepository,
) domain.TurnosNTLISUseCase {

	return &turnosNTLISApp{
		repository: repository,
	}
}

func (a *turnosNTLISApp) BuscarSedesNTService(
	ctx context.Context,
) (domain.SedesNTResponse, error) {

	sedes, err := a.repository.GetSedes(ctx)
	if err != nil {
		return domain.SedesNTResponse{}, err
	}

	response := domain.SedesNTResponse{
		Status: 200,
		Data:   sedes,
	}

	return response, nil
}
