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

func (a *turnosNTLISApp) BuscarSedesNTService(ctx context.Context) (domain.SedesNTResponse, error) {

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

func (a *turnosNTLISApp) BuscarServiciosNTXSedeService(
	ctx context.Context,
	sedeID int,
) (domain.ServiciosNTXSedeResponse, error) {

	servicios, err := a.repository.GetServiciosNTXSede(ctx, sedeID)
	if err != nil {
		return domain.ServiciosNTXSedeResponse{}, err
	}

	response := domain.ServiciosNTXSedeResponse{
		Status: 200,
		Data:   servicios,
	}

	return response, nil
}

func (a *turnosNTLISApp) BuscarTaquillasNTService(ctx context.Context) (domain.TaquillaNTResponse, error) {

	taquillas, err := a.repository.GetTaquillasNT(ctx)
	if err != nil {
		return domain.TaquillaNTResponse{}, err
	}

	response := domain.TaquillaNTResponse{
		Status: 200,
		Data:   taquillas,
	}

	return response, nil
}
