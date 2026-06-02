package domain

import "context"

type TurnosNTLISUseCase interface {
	BuscarSedesNTService(context.Context) (SedesNTResponse, error)
	BuscarServiciosNTXSedeService(context.Context, int) (ServiciosNTXSedeResponse, error)
	BuscarTaquillasNTService(context.Context) (TaquillaNTResponse, error)
	BuscarTaquillasXServicio(context.Context, int, int) (TaquillaxSedexServicioResponse, error)
	BuscarMotivosTaquilla(context.Context) (TaquillaMotivoDescansoResponse, error)
	ActualizarEstadoTaquilla(context.Context, TaquillasEstadoRequest) (TaquillasEstadoResponse, error)
	ActualizarEstadoAtencion(context.Context, AtencionEstadoRequest) (AtencionEstadoResponse, error)
	BuscarMotivosCancelacion(context.Context) (Response, error)
	BuscarServiciosDisponiblesParaTransferenciaService(context.Context, int, int, int) (ResponseTransfer, error)
	LlamadoTurnoService(context.Context, int, int) (LlamadoTurnoResponse, error)
	LlamadoTurnoPostService(context.Context, LlamadoTurnoPostRequest) (LlamadoTurnoPostResponse, error)
}

type TurnosNTLISRepository interface {
	GetSedes(ctx context.Context) ([]SedeNT, error)
	GetServiciosNTXSede(ctx context.Context, sedeID int) ([]ServicioNT, error)
	GetTaquillasNT(ctx context.Context) ([]TaquillaNT, error)
	GetTaquillasXServicio(ctx context.Context, sedeID int, servicioID int) ([]TaquillaxSedexServicioData, error)
	GetMotivosTaquilla(ctx context.Context) ([]MotivoDescanso, error)
	UpdateEstadoTaquilla(ctx context.Context, sedeID int, taquillaID int, estado int) error
	UpdateEstadoAtencion(ctx context.Context, idTurno int, idServicio int, idTaquilla int, estado int) error
	GetMotivosCancelacion(ctx context.Context) ([]Reason, error)
	GetServiciosDisponiblesParaTransferencia(ctx context.Context, idSede int, idServicio int, idTurno int) ([]TransferReason, error)
	GetLlamadoTurno(ctx context.Context, turno int, idServicio int) (bool, error)
	CallTurnPost(ctx context.Context, req LlamadoTurnoPostRequest) (LlamadoTurnoPostData, error)
}
