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

func (a *CreacionTurnosApp) BuscarTipoDocumentoService(ctx context.Context) (domain.TipoDocumentosesResponse, error) {
	return a.repository.BuscarTipoDocumento(ctx)
}

func (a *CreacionTurnosApp) BuscarCompaniaService(ctx context.Context, req domain.CompaniaRequest) (domain.CompaniasResponse, error) {
	return a.repository.BuscarCompania(ctx, req)
}

func (a *CreacionTurnosApp) VerificarConfigCompaniasService(ctx context.Context) (domain.ConfigCompaniasResponse, error) {
	return a.repository.VerificarConfigCompanias(ctx)
}

func (a *CreacionTurnosApp) BuscarTipoServicioService(ctx context.Context, req domain.TipoServicioRequest) (domain.TipoServiciosResponse, error) {
	return a.repository.BuscarTipoServicio(ctx, req)
}

func (a *CreacionTurnosApp) BuscarModuloService(ctx context.Context, req domain.ModuloRequest) (domain.ModulosResponse, error) {
	return a.repository.BuscarModulo(ctx, req)
}

func (a *CreacionTurnosApp) BuscarSedesService(ctx context.Context) (domain.SedesResponse, error) {
	return a.repository.BuscarSedes(ctx)
}

func (a *CreacionTurnosApp) BuscarTipoTurnoService(ctx context.Context, req domain.TipoTurnoRequest) (domain.TipoTurnosResponse, error) {
	return a.repository.BuscarTipoTurno(ctx, req)
}
