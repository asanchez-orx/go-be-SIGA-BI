package domain

import "context"

type CreacionTurnosUseCase interface {
	BuscarCreacionTurnosService(context.Context, CreacionTurnosRequest) (CreacionTurnosesResponse, error)
	CrearCreacionTurnosService(context.Context, CreacionTurnosRequest) error
	BuscarTipoDocumentoService(context.Context) (TipoDocumentosesResponse, error)
	BuscarCompaniaService(context.Context, CompaniaRequest) (CompaniasResponse, error)
	VerificarConfigCompaniasService(context.Context) (ConfigCompaniasResponse, error)
	BuscarTipoServicioService(context.Context, TipoServicioRequest) (TipoServiciosResponse, error)
	BuscarModuloService(context.Context, ModuloRequest) (ModulosResponse, error)
	BuscarSedesService(context.Context) (SedesResponse, error)
	BuscarTipoTurnoService(context.Context, TipoTurnoRequest) (TipoTurnosResponse, error)
}

type CreacionTurnosRepository interface {
	BuscarCreacionTurnos(context.Context, CreacionTurnosRequest) (CreacionTurnosesResponse, error)
	CrearCreacionTurnos(context.Context, CreacionTurnosRequest) error
	BuscarTipoDocumento(context.Context) (TipoDocumentosesResponse, error)
	BuscarCompania(context.Context, CompaniaRequest) (CompaniasResponse, error)
	VerificarConfigCompanias(context.Context) (ConfigCompaniasResponse, error)
	BuscarTipoServicio(context.Context, TipoServicioRequest) (TipoServiciosResponse, error)
	BuscarModulo(context.Context, ModuloRequest) (ModulosResponse, error)
	BuscarSedes(context.Context) (SedesResponse, error)
	BuscarTipoTurno(context.Context, TipoTurnoRequest) (TipoTurnosResponse, error)
}
