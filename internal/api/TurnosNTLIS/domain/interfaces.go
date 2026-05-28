package domain

import "context"

type TurnosNTLISUseCase interface {
	BuscarSedesNTService(context.Context) (SedesNTResponse, error)
}

type TurnosNTLISRepository interface {
	GetSedes(ctx context.Context) ([]SedeNT, error)
}
