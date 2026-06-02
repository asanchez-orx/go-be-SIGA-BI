package mssql

import (
	"context"
	"fmt"
	"strconv"

	"develop.private/CLTech/besigabi/internal/api/TurnosNTLIS/domain"
	"develop.private/CLTech/vulcano/infra/database"
)

type TurnosNTLISRepo struct {
	db database.Database
}

func NewTurnosNTLISRepo(db database.Database) *TurnosNTLISRepo {
	return &TurnosNTLISRepo{
		db: db,
	}
}

func (r *TurnosNTLISRepo) GetSedes(ctx context.Context) ([]domain.SedeNT, error) {

	rows, err := r.db.Query(ctx, qrySedes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sedes []domain.SedeNT

	for rows.Next() {

		var sede domain.SedeNT

		err := rows.Scan(
			&sede.ID,
			&sede.Code,
			&sede.Name,
			&sede.Description,
			&sede.RegisterDate,
			&sede.State,
		)

		if err != nil {
			return nil, err
		}

		sedes = append(sedes, sede)
	}

	return sedes, nil
}

func (r *TurnosNTLISRepo) GetServiciosNTXSede(ctx context.Context, sedeID int) ([]domain.ServicioNT, error) {

	rows, err := r.db.Query(ctx, qryServiciosNTXSede, sedeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servicios []domain.ServicioNT

	for rows.Next() {

		var servicio domain.ServicioNT

		err := rows.Scan(
			&servicio.ID,
			&servicio.Code,
			&servicio.Name,
			&servicio.Description,
			&servicio.RegisterDate,
		)

		if err != nil {
			return nil, err
		}

		servicios = append(servicios, servicio)
	}

	return servicios, nil
}

func (r *TurnosNTLISRepo) GetTaquillasNT(ctx context.Context) ([]domain.TaquillaNT, error) {

	rows, err := r.db.Query(ctx, queryTaquillas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taquillas []domain.TaquillaNT

	for rows.Next() {

		var taquilla domain.TaquillaNT

		err := rows.Scan(
			&taquilla.ID,
			&taquilla.Code,
			&taquilla.Name,
			&taquilla.State,
		)

		if err != nil {
			return nil, err
		}

		taquillas = append(taquillas, taquilla)
	}

	return taquillas, nil
}

func (r *TurnosNTLISRepo) GetTaquillasXServicio(ctx context.Context, sedeID int, servicioID int) ([]domain.TaquillaxSedexServicioData, error) {

	rows, err := r.db.Query(ctx, queryTaquillasxSedeServicio, sedeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taquillas []domain.TaquillaxSedexServicioData

	for rows.Next() {

		var taquilla domain.TaquillaxSedexServicioData

		err := rows.Scan(
			&taquilla.Id,
			&taquilla.Name,
		)

		if err != nil {
			return nil, err
		}

		taquillas = append(taquillas, taquilla)
	}

	return taquillas, nil
}

func (r *TurnosNTLISRepo) GetMotivosTaquilla(ctx context.Context) ([]domain.MotivoDescanso, error) {

	rows, err := r.db.Query(ctx, queryMotivosDescanso)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var motivos []domain.MotivoDescanso

	for rows.Next() {

		var motivo domain.MotivoDescanso

		err := rows.Scan(
			&motivo.Id,
			&motivo.Name,
			&motivo.Description,
			&motivo.RegisterDate,
			&motivo.Type,
			&motivo.State,
		)

		if err != nil {
			return nil, err
		}

		motivos = append(motivos, motivo)
	}

	return motivos, nil

}

func (r *TurnosNTLISRepo) UpdateEstadoTaquilla(ctx context.Context, sedeID int, taquillaID int, estado int) error {
	rows, err := r.db.Exec(ctx, qryUpdateEstadoTaquilla, estado, sedeID, taquillaID)
	fmt.Printf("UpdateEstadoTaquilla: sedeID=%d, taquillaID=%d, estado=%d -> rows affected: %d\n", sedeID, taquillaID, estado, rows)
	if err != nil {
		fmt.Printf("Error in UpdateEstadoTaquilla: %v\n", err)
	}
	return err
}

func (r *TurnosNTLISRepo) UpdateEstadoAtencion(ctx context.Context, idTurno int, idServicio int, idTaquilla int, estado int) error {
	rows, err := r.db.Exec(ctx, qryUpdateEstadoAtencion, estado, idTurno, idServicio, idTaquilla)
	fmt.Printf("UpdateEstadoAtencion: idTurno=%d, idServicio=%d, idTaquilla=%d, estado=%d -> rows affected: %d\n", idTurno, idServicio, idTaquilla, estado, rows)
	if err != nil {
		fmt.Printf("Error in UpdateEstadoAtencion: %v\n", err)
	}
	return err
}

func (r *TurnosNTLISRepo) GetMotivosCancelacion(ctx context.Context) ([]domain.Reason, error) {

	rows, err := r.db.Query(ctx, qryMotivosCancelacion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var motivos []domain.Reason

	for rows.Next() {

		var motivo domain.Reason

		err := rows.Scan(
			&motivo.ID,
			&motivo.Name,
			&motivo.Type,
			&motivo.State,
		)

		if err != nil {
			return nil, err
		}

		motivos = append(motivos, motivo)
	}

	return motivos, nil

}

func (r *TurnosNTLISRepo) GetServiciosDisponiblesParaTransferencia(ctx context.Context, idSede int, idServicio int, idTurno int) ([]domain.TransferReason, error) {

	rowsOri, err := r.db.Query(ctx, qryServicioOrigen, idTurno)
	if err != nil {
		return nil, err
	}
	var serviceOriID int
	if rowsOri.Next() {
		err = rowsOri.Scan(&serviceOriID)
		if err != nil {
			rowsOri.Close()
			return nil, err
		}
	}
	rowsOri.Close()

	rows, err := r.db.Query(ctx, qryServiciosDisponiblesParaTransferencia, idServicio, idSede)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []domain.TransferReason

	for rows.Next() {
		var tr domain.TransferReason
		var sedeID int

		err := rows.Scan(
			&sedeID,
			&tr.ServiceDes.ID,
			&tr.ServiceDes.Code,
			&tr.ServiceDes.Name,
			&tr.WaitTime,
		)

		if err != nil {
			return nil, err
		}

		tr.ID = 0
		tr.Branch.ID = sedeID
		tr.ServiceOri.ID = serviceOriID
		tr.PendingTurns = 0
		tr.Enabled = true

		transfers = append(transfers, tr)
	}

	return transfers, nil
}

func (r *TurnosNTLISRepo) GetLlamadoTurno(ctx context.Context, turno int, idServicio int) (bool, error) {
	rows, err := r.db.Query(ctx, qryLlamadoTurno, turno, idServicio)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (r *TurnosNTLISRepo) CallTurnPost(ctx context.Context, req domain.LlamadoTurnoPostRequest) (domain.LlamadoTurnoPostData, error) {
	_, err := r.db.Exec(ctx, qryUpdateTurnoPost, req.IdTurn, req.Point.Branch.ID, req.Point.Service.ID)
	if err != nil {
		return domain.LlamadoTurnoPostData{}, err
	}

	var data domain.LlamadoTurnoPostData
	rows, err := r.db.Query(ctx, qrySelectTurnoPost, req.IdTurn, req.Point.Branch.ID, req.Point.Service.ID)
	if err != nil {
		return domain.LlamadoTurnoPostData{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var pid string
		var ppid string
		err := rows.Scan(
			&data.ID,
			&data.Number,
			&data.TurnType.ID,
			&data.TurnType.Code,
			&data.TurnType.Name,
			&pid,
			&ppid,
			&data.Service.ID,
			&data.Service.Name,
		)
		if err != nil {
			return domain.LlamadoTurnoPostData{}, err
		}
		
		importStrconv := false // Just for reference if we need to convert
		if pid != "" {
			var errConv error
			if data.Patient.ID, errConv = strconv.Atoi(pid); errConv != nil {
				data.Patient.ID = 0
			}
			_ = importStrconv
		}
		data.Patient.PatientID = ppid
		
		data.Priority = 1
		data.State = 1
		data.Attended = false
		data.Transferible = true
		data.Finalizable = true

		return data, nil
	}

	return domain.LlamadoTurnoPostData{}, fmt.Errorf("turn not found")
}

