package mssql

import (
	"context"
	"encoding/json"
	"time"

	"develop.private/CLTech/besigabi/internal/api/SKL/domain"
	"develop.private/CLTech/vulcano/infra/database"
)

type SKLRepo struct {
	db database.Database
}

func NewSKLRepo(db database.Database) *SKLRepo {
	return &SKLRepo{
		db: db,
	}
}

func (r *SKLRepo) GetTaquillas(ctx context.Context, req domain.TaquillasRequest) (domain.TaquillasResponse, error) {
	rows, err := r.db.Query(ctx, qryTaquillas, req.IdSede, req.CodModulo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.TaquillasResponse{}
	for rows.Next() {
		item := domain.TaquillaResponse{}
		if err := rows.Scan(&item.IdTaquilla, &item.CodTaquilla, &item.NomTaquilla, &item.EstadoTaquilla, &item.NIdModulo); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *SKLRepo) GetServiciosSiga(ctx context.Context, req domain.ServiciosSigaRequest) (domain.ServiciosSigaResponse, error) {
	rows, err := r.db.Query(ctx, qryServiciosSiga, req.IdSede, req.CodModulo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.ServiciosSigaResponse{}
	for rows.Next() {
		item := domain.ServicioSigaResponse{}
		if err := rows.Scan(&item.IdServicio, &item.CodServicio, &item.NomServicio); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *SKLRepo) GetTurnosDisponibles(ctx context.Context, req domain.TurnosDisponiblesRequest) (domain.TurnosDisponiblesResponse, error) {

	// =========================================
	// 1. PRIORIDADES POR TAQUILLA
	// =========================================

	prioridadesMap := make(map[int]int)

	rowsPrioridades, err := r.db.Query(
		ctx,
		qryPrioridadesPorTaquilla,
		req.IdTaquilla,
	)

	if err != nil {
		return nil, err
	}

	defer rowsPrioridades.Close()

	for rowsPrioridades.Next() {

		var (
			idTipoTurno int
			nPrioridad  int
		)

		if err := rowsPrioridades.Scan(
			&idTipoTurno,
			&nPrioridad,
		); err != nil {
			return nil, err
		}

		prioridadesMap[idTipoTurno] = nPrioridad
	}

	// =========================================
	// 2. TIEMPOS LIMITE POR SEDE
	// =========================================

	tiemposMap := make(map[int]int)

	rowsTiempos, err := r.db.Query(
		ctx,
		qryTiemposPorSede,
		req.IdSede,
	)

	if err != nil {
		return nil, err
	}

	defer rowsTiempos.Close()

	for rowsTiempos.Next() {

		var (
			idTipoTurno int
			nCantidad   int
		)

		if err := rowsTiempos.Scan(
			&idTipoTurno,
			&nCantidad,
		); err != nil {
			return nil, err
		}

		tiemposMap[idTipoTurno] = nCantidad
	}

	// =========================================
	// 3. CONSULTAR TURNOS
	// =========================================

	rows, err := r.db.Query(
		ctx,
		qryTurnosDisponibles,
		req.IdServicio,
		req.IdSede,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := domain.TurnosDisponiblesResponse{}

	for rows.Next() {

		var (
			idSede          int
			nombreSede      string
			idServicio      int
			codServicio     string
			nomServicio     string
			idTipoTurno     int
			codTipoTurno    string
			nomTipoTurno    string
			fechaHoraTurno  string
			idTipoDocumento int
			idEmpresa       int
			codEmpresa      string
			nomEmpresa      string
			numeroTurno     string
			numeroOrden     string
			jsonPaciente    string
		)

		if err := rows.Scan(
			&idSede,
			&nombreSede,
			&idServicio,
			&codServicio,
			&nomServicio,
			&idTipoTurno,
			&codTipoTurno,
			&nomTipoTurno,
			&fechaHoraTurno,
			&idTipoDocumento,
			&idEmpresa,
			&codEmpresa,
			&nomEmpresa,
			&numeroTurno,
			&numeroOrden,
			&jsonPaciente,
		); err != nil {
			return nil, err
		}

		// =========================================
		// DESERIALIZAR PACIENTE
		// =========================================

		var paciente domain.PacienteTurno

		if jsonPaciente != "" {

			if err := json.Unmarshal(
				[]byte(jsonPaciente),
				&paciente,
			); err != nil {

				// NO rompe flujo
				paciente = domain.PacienteTurno{}
			}
		}

		// =========================================
		// CALCULAR MINUTOS ESPERA
		// =========================================

		minutosEspera := 0

		if fechaHoraTurno != "" &&
			len(fechaHoraTurno) == 14 {

			sHoraTurno := fechaHoraTurno[8:14]

			tHoraTurno, err := time.Parse(
				"150405",
				sHoraTurno,
			)

			if err == nil {

				now := time.Now()

				minutosTurno :=
					tHoraTurno.Hour()*60 +
						tHoraTurno.Minute()

				minutosAhora :=
					now.Hour()*60 +
						now.Minute()

				minutosEspera =
					minutosAhora - minutosTurno

				if minutosEspera < 0 {
					minutosEspera = 0
				}
			}
		}

		// =========================================
		// CALCULAR SEMAFORO TIEMPO
		// =========================================

		tiempo := 2

		if nCantidad, found := tiemposMap[idTipoTurno]; found {

			if minutosEspera > nCantidad {
				tiempo = 3
			}
		}

		// =========================================
		// OBTENER PRIORIDAD
		// =========================================

		prioridad := prioridadesMap[idTipoTurno]

		// =========================================
		// ARMAR RESPONSE
		// =========================================

		res = append(res, domain.TurnoDisponibleResponse{
			IdSede:          idSede,
			NombreSede:      nombreSede,
			IdServicio:      idServicio,
			CodServicio:     codServicio,
			NomServicio:     nomServicio,
			IdTipoTurno:     idTipoTurno,
			CodTipoTurno:    codTipoTurno,
			NomTipoTurno:    nomTipoTurno,
			NumeroTurno:     numeroTurno,
			FechaHoraTurno:  fechaHoraTurno,
			IdTipoDocumento: idTipoDocumento,
			IdEmpresa:       idEmpresa,
			CodEmpresa:      codEmpresa,
			NomEmpresa:      nomEmpresa,
			NumeroOrden:     numeroOrden,
			MinutosEspera:   minutosEspera,
			Prioridad:       prioridad,
			Tiempo:          tiempo,
			Paciente:        paciente,
		})
	}

	return res, nil
}

func (r *SKLRepo) GetSedesUsuario(ctx context.Context, req domain.SedesUsuarioRequest) (domain.SedesUsuarioResponse, error) {
	rows, err := r.db.Query(ctx, qrySedesUsuario, req.Usuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.SedesUsuarioResponse{}
	for rows.Next() {
		item := domain.SedeUsuarioResponse{}
		if err := rows.Scan(&item.IdSede, &item.CodSede, &item.NomSede); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *SKLRepo) ConsumirCredenciales(ctx context.Context, req domain.ConsumirCredencialesRequest) (domain.ConsumirCredencialesResponse, error) {
	var res domain.ConsumirCredencialesResponse
	err := r.db.QueryRow(ctx, qryConsumirCredenciales, req.Usuario, req.Contrasena).Scan(&res.IdUsuario)
	if err != nil {
		return res, err
	}
	return res, nil
}


func (r *SKLRepo) GetTurnosDisponiblesConOrden(ctx context.Context, req domain.TurnosDisponiblesRequest) (domain.TurnosDisponiblesResponse, error) {

	// =========================================
	// 1. PRIORIDADES POR TAQUILLA
	// =========================================

	prioridadesMap := make(map[int]int)

	rowsPrioridades, err := r.db.Query(
		ctx,
		qryPrioridadesPorTaquilla,
		req.IdTaquilla,
	)

	if err != nil {
		return nil, err
	}

	defer rowsPrioridades.Close()

	for rowsPrioridades.Next() {

		var (
			idTipoTurno int
			nPrioridad  int
		)

		if err := rowsPrioridades.Scan(
			&idTipoTurno,
			&nPrioridad,
		); err != nil {
			return nil, err
		}

		prioridadesMap[idTipoTurno] = nPrioridad
	}

	// =========================================
	// 2. TIEMPOS LIMITE POR SEDE
	// =========================================

	tiemposMap := make(map[int]int)

	rowsTiempos, err := r.db.Query(
		ctx,
		qryTiemposPorSede,
		req.IdSede,
	)

	if err != nil {
		return nil, err
	}

	defer rowsTiempos.Close()

	for rowsTiempos.Next() {

		var (
			idTipoTurno int
			nCantidad   int
		)

		if err := rowsTiempos.Scan(
			&idTipoTurno,
			&nCantidad,
		); err != nil {
			return nil, err
		}

		tiemposMap[idTipoTurno] = nCantidad
	}

	// =========================================
	// 3. CONSULTAR TURNOS
	// =========================================

	rows, err := r.db.Query(
		ctx,
		qryTurnosDisponiblesConOrden,
		req.IdServicio,
		req.IdSede,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := domain.TurnosDisponiblesResponse{}

	for rows.Next() {

		var (
			idSede          int
			nombreSede      string
			idServicio      int
			codServicio     string
			nomServicio     string
			idTipoTurno     int
			codTipoTurno    string
			nomTipoTurno    string
			fechaHoraTurno  string
			idTipoDocumento int
			idEmpresa       int
			codEmpresa      string
			nomEmpresa      string
			numeroTurno     string
			numeroOrden     string
			jsonPaciente    string
		)

		if err := rows.Scan(
			&idSede,
			&nombreSede,
			&idServicio,
			&codServicio,
			&nomServicio,
			&idTipoTurno,
			&codTipoTurno,
			&nomTipoTurno,
			&fechaHoraTurno,
			&idTipoDocumento,
			&idEmpresa,
			&codEmpresa,
			&nomEmpresa,
			&numeroTurno,
			&numeroOrden,
			&jsonPaciente,
		); err != nil {
			return nil, err
		}

		// =========================================
		// DESERIALIZAR PACIENTE
		// =========================================

		var paciente domain.PacienteTurno

		if jsonPaciente != "" {

			if err := json.Unmarshal(
				[]byte(jsonPaciente),
				&paciente,
			); err != nil {

				// NO rompe flujo
				paciente = domain.PacienteTurno{}
			}
		}

		// =========================================
		// CALCULAR MINUTOS ESPERA
		// =========================================

		minutosEspera := 0

		if fechaHoraTurno != "" &&
			len(fechaHoraTurno) == 14 {

			sHoraTurno := fechaHoraTurno[8:14]

			tHoraTurno, err := time.Parse(
				"150405",
				sHoraTurno,
			)

			if err == nil {

				now := time.Now()

				minutosTurno :=
					tHoraTurno.Hour()*60 +
						tHoraTurno.Minute()

				minutosAhora :=
					now.Hour()*60 +
						now.Minute()

				minutosEspera =
					minutosAhora - minutosTurno

				if minutosEspera < 0 {
					minutosEspera = 0
				}
			}
		}

		// =========================================
		// CALCULAR SEMAFORO TIEMPO
		// =========================================

		tiempo := 2

		if nCantidad, found := tiemposMap[idTipoTurno]; found {

			if minutosEspera > nCantidad {
				tiempo = 3
			}
		}

		// =========================================
		// OBTENER PRIORIDAD
		// =========================================

		prioridad := prioridadesMap[idTipoTurno]

		// =========================================
		// ARMAR RESPONSE
		// =========================================

		res = append(res, domain.TurnoDisponibleResponse{
			IdSede:          idSede,
			NombreSede:      nombreSede,
			IdServicio:      idServicio,
			CodServicio:     codServicio,
			NomServicio:     nomServicio,
			IdTipoTurno:     idTipoTurno,
			CodTipoTurno:    codTipoTurno,
			NomTipoTurno:    nomTipoTurno,
			NumeroTurno:     numeroTurno,
			FechaHoraTurno:  fechaHoraTurno,
			IdTipoDocumento: idTipoDocumento,
			IdEmpresa:       idEmpresa,
			CodEmpresa:      codEmpresa,
			NomEmpresa:      nomEmpresa,
			NumeroOrden:     numeroOrden,
			MinutosEspera:   minutosEspera,
			Prioridad:       prioridad,
			Tiempo:          tiempo,
			Paciente:        paciente,
		})
	}

	return res, nil
}

