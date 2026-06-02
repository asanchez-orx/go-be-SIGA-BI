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

func (a *turnosNTLISApp) BuscarTaquillasXServicio(
	ctx context.Context,
	sedeID int,
	servicioID int,
) (domain.TaquillaxSedexServicioResponse, error) {

	taquillas, err := a.repository.GetTaquillasXServicio(ctx, sedeID, servicioID)
	if err != nil {
		return domain.TaquillaxSedexServicioResponse{}, err
	}

	response := domain.TaquillaxSedexServicioResponse{
		Status: 200,
		Data:   taquillas,
	}

	return response, nil
}

func (a *turnosNTLISApp) BuscarMotivosTaquilla(ctx context.Context) (domain.TaquillaMotivoDescansoResponse, error) {

	motivos, err := a.repository.GetMotivosTaquilla(ctx)
	if err != nil {
		return domain.TaquillaMotivoDescansoResponse{}, err
	}

	response := domain.TaquillaMotivoDescansoResponse{
		Status: 200,
		Data:   motivos,
	}

	return response, nil
}

func (a *turnosNTLISApp) ActualizarEstadoTaquilla(ctx context.Context, req domain.TaquillasEstadoRequest) (domain.TaquillasEstadoResponse, error) {

	updateState := 0
	if req.Action == 1 {
		updateState = 1
	} else if req.Action == 2 {
		updateState = 2
	} else if req.Action == 4 {
		updateState = 0
	} else {
		updateState = req.Action
	}

	err := a.repository.UpdateEstadoTaquilla(ctx, req.Branch.ID, req.Point.ID, updateState)
	if err != nil {
		return domain.TaquillasEstadoResponse{}, err
	}

	respData := req
	respData.Branch.State = updateState

	response := domain.TaquillasEstadoResponse{
		Status: 200,
		Data:   respData,
	}

	return response, nil
}

func (a *turnosNTLISApp) ActualizarEstadoAtencion(ctx context.Context, req domain.AtencionEstadoRequest) (domain.AtencionEstadoResponse, error) {

	updateState := req.State
	if req.State == 3 {
		updateState = 2
	} else if req.State == 4 {
		updateState = 5
	} else if req.State == 5 {
		updateState = 6
	} else if req.State == 6 {
		updateState = 5
	}

	err := a.repository.UpdateEstadoAtencion(ctx, req.Turn.ID, req.Service.ID, req.PointOfCare.ID, updateState)
	if err != nil {
		return domain.AtencionEstadoResponse{}, err
	}

	response := domain.AtencionEstadoResponse{
		Status: 200,
		Data:   req.Service.ID,
	}

	return response, nil
}

func (a *turnosNTLISApp) BuscarMotivosCancelacion(ctx context.Context) (domain.Response, error) {

	motivos, err := a.repository.GetMotivosCancelacion(ctx)
	if err != nil {
		return domain.Response{}, err
	}

	response := domain.Response{
		Status: 200,
		Data:   motivos,
	}

	return response, nil
}

func (a *turnosNTLISApp) BuscarServiciosDisponiblesParaTransferenciaService(ctx context.Context, idSede int, idServicio int, idTurno int) (domain.ResponseTransfer, error) {

	servicios, err := a.repository.GetServiciosDisponiblesParaTransferencia(ctx, idSede, idServicio, idTurno)
	if err != nil {
		return domain.ResponseTransfer{}, err
	}

	response := domain.ResponseTransfer{
		Status: 200,
		Data:   servicios,
	}

	return response, nil
}

func (a *turnosNTLISApp) LlamadoTurnoService(ctx context.Context, turno int, idServicio int) (domain.LlamadoTurnoResponse, error) {

	turnoExists, err := a.repository.GetLlamadoTurno(ctx, turno, idServicio)
	if err != nil {
		return domain.LlamadoTurnoResponse{}, err
	}

	response := domain.LlamadoTurnoResponse{
		Status: 200,
		Data:   turnoExists,
	}

	return response, nil
}

func (a *turnosNTLISApp) LlamadoTurnoPostService(ctx context.Context, req domain.LlamadoTurnoPostRequest) (domain.LlamadoTurnoPostResponse, error) {
	data, err := a.repository.CallTurnPost(ctx, req)
	if err != nil {
		return domain.LlamadoTurnoPostResponse{}, err
	}

	return domain.LlamadoTurnoPostResponse{
		Status: 200,
		Data:   data,
	}, nil
}
