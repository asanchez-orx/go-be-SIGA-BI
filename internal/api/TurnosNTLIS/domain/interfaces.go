package domain

import "context"

type TurnosNTLISUseCase interface {
	BuscarSedesNTService(context.Context) (SedesNTResponse, error)
	BuscarServiciosNTXSedeService(context.Context, int) (ServiciosNTXSedeResponse, error)
	BuscarTaquillasNTService(context.Context) (TaquillaNTResponse, error)
}

type TurnosNTLISRepository interface {
	GetSedes(ctx context.Context) ([]SedeNT, error)
	GetServiciosNTXSede(ctx context.Context, sedeID int) ([]ServicioNT, error)
	GetTaquillasNT(ctx context.Context) ([]TaquillaNT, error)
}
