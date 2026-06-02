package mssql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"develop.private/CLTech/besigabi/internal/api/creacionTurnos/domain"
	"develop.private/CLTech/vulcano/infra/database"
)

type CreacionTurnosRepo struct {
	db database.Database
}

func NewCreacionTurnosRepo(db database.Database) *CreacionTurnosRepo {
	return &CreacionTurnosRepo{
		db: db,
	}
}

func (r *CreacionTurnosRepo) BuscarCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) (domain.CreacionTurnosesResponse, error) {
	return domain.CreacionTurnosesResponse{}, nil
}

func (r *CreacionTurnosRepo) CrearCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) error {
	return nil
}

func (r *CreacionTurnosRepo) BuscarTipoDocumento(ctx context.Context) (domain.TipoDocumentosesResponse, error) {
	rows, err := r.db.Query(ctx, qryTipoDocumento)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.TipoDocumentosesResponse{}
	for rows.Next() {
		item := domain.TipoDocumentoResponse{}
		if err := rows.Scan(&item.IdTipoDoc, &item.CodTipoDoc, &item.NomTipoDoc); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarCompania(ctx context.Context, req domain.CompaniaRequest) (domain.CompaniasResponse, error) {
	var rows database.Rows
	var err error

	if req.IdSede == -1 {
		rows, err = r.db.Query(ctx, qryCompanias)
	} else {
		rows, err = r.db.Query(ctx, qryCompaniasPorSede, req.IdSede)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.CompaniasResponse{}
	for rows.Next() {
		item := domain.CompaniaResponse{}
		if err := rows.Scan(&item.IdCompania, &item.CodigoCompania, &item.NombreCompania); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) VerificarConfigCompanias(ctx context.Context) (domain.ConfigCompaniasResponse, error) {
	var res domain.ConfigCompaniasResponse

	err := r.db.QueryRow(ctx, qryVerificarConfigCompanias).Scan(&res.ManejaEmpresas)
	if err != nil {
		if database.NoRows(err) {
			return domain.ConfigCompaniasResponse{ManejaEmpresas: 0}, nil
		}
		return domain.ConfigCompaniasResponse{}, err
	}

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarTipoServicio(ctx context.Context, req domain.TipoServicioRequest) (domain.TipoServiciosResponse, error) {
	var rows database.Rows
	var err error

	if req.IdCompania == -1 {
		rows, err = r.db.Query(ctx, qryTipoServicio, string(req.CodigoModulo), req.IdSede)
	} else {
		rows, err = r.db.Query(ctx, qryTipoServicioPorCompania, req.IdCompania, string(req.CodigoModulo), req.IdSede)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.TipoServiciosResponse{}
	for rows.Next() {
		item := domain.TipoServicioResponse{}
		if err := rows.Scan(&item.IdServicio, &item.CodigoServicio, &item.NombreServicio); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarModulo(ctx context.Context, req domain.ModuloRequest) (domain.ModulosResponse, error) {
	rows, err := r.db.Query(ctx, qryModulos, req.IdSede)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.ModulosResponse{}
	for rows.Next() {
		var itemDB domain.ModuloDB
		if err := rows.Scan(&itemDB.Id, &itemDB.IdSede, &itemDB.JsonModulos); err != nil {
			return nil, err
		}

		var modulos []domain.ModuloResponse
		if err := json.Unmarshal([]byte(itemDB.JsonModulos), &modulos); err != nil {
			continue
		}
		res = append(res, modulos...)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarSedes(ctx context.Context) (domain.SedesResponse, error) {
	rows, err := r.db.Query(ctx, qrySedes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.SedesResponse{}
	for rows.Next() {
		item := domain.SedeResponse{}
		if err := rows.Scan(&item.IdSede, &item.NomSede, &item.CodSede); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarTipoTurno(ctx context.Context, req domain.TipoTurnoRequest) (domain.TipoTurnosResponse, error) {
	rows, err := r.db.Query(ctx, qryTipoTurno, req.IdServicio, req.IdSede)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := domain.TipoTurnosResponse{}
	for rows.Next() {
		item := domain.TipoTurnoResponse{}
		if err := rows.Scan(&item.IdTipoTurno, &item.CodTipoTurno, &item.NomTipoTurno); err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	return res, nil
}

func (r *CreacionTurnosRepo) ConfirmarConfigSedes(ctx context.Context) (interface{}, error) {
	var datConfig string

	err := r.db.QueryRow(ctx, qryConfirmarConfigSedes).Scan(&datConfig)
	if err != nil && !database.NoRows(err) {
		return nil, err
	}

	if datConfig == "1" {
		return 1, nil
	}

	var sede domain.SedeResponse
	err = r.db.QueryRow(ctx, qrySedeDefault).Scan(&sede.IdSede, &sede.CodSede, &sede.NomSede)
	if err != nil {
		if database.NoRows(err) {
			// No se encontró sede por default
			return nil, nil
		}
		return nil, err
	}

	return sede, nil
}

func (r *CreacionTurnosRepo) CargarConfigLIS(ctx context.Context) (domain.ConfigLISResponse, error) {
	var res domain.ConfigLISResponse
	if err := r.db.QueryRow(ctx, qryCargarConfigLIS).Scan(&res.SeparadorMuestra); err != nil {
		if database.NoRows(err) {
			return domain.ConfigLISResponse{}, fmt.Errorf("configuración LIS no encontrada")
		}
		return domain.ConfigLISResponse{}, err
	}

	return res, nil
}

// /////====================//
// /////   Creacion Turnos   //
// ///////////////////////////
// Inicio del servicio que crea turnos: flujo handler -> app -> repo
// Aquí comienza la implementación de la lógica para generar e insertar
// un nuevo turno en la tabla LAB5843.
func (r *CreacionTurnosRepo) CrearTurno(ctx context.Context, req domain.CrearTurnoRequest) (domain.CrearTurnoResponse, error) {
	if req.Accion != "Nuevo" {
		return domain.CrearTurnoResponse{}, fmt.Errorf("acción no soportada: %s", req.Accion)
	}

	// Buscar sede por ID
	var sede domain.SedeResponse
	if err := r.db.QueryRow(ctx, qrySedeById, req.IdSede).Scan(&sede.IdSede, &sede.CodSede, &sede.NomSede); err != nil {
		return domain.CrearTurnoResponse{}, domain.ErrSedeNotFound
	}

	// Buscar compañía por ID sólo si fue especificada
	var compania domain.CompaniaResponse
	if req.IdCompania != -1 {
		if err := r.db.QueryRow(ctx, qryCompaniaById, req.IdCompania).Scan(&compania.IdCompania, &compania.CodigoCompania, &compania.NombreCompania); err != nil {
			return domain.CrearTurnoResponse{}, domain.ErrCompaniaNotFound
		}
	}

	// Buscar tipo de turno por ID
	var tipoTurno domain.TipoTurnoResponse
	if err := r.db.QueryRow(ctx, qryTipoTurnoById, req.IdTipoTurno).Scan(&tipoTurno.IdTipoTurno, &tipoTurno.CodTipoTurno, &tipoTurno.NomTipoTurno); err != nil {
		return domain.CrearTurnoResponse{}, domain.ErrTipoTurnoNotFound
	}

	// Buscar servicio por ID
	var servicio domain.TipoServicioResponse
	if err := r.db.QueryRow(ctx, qryServicioById, req.IdServicio).Scan(&servicio.IdServicio, &servicio.CodigoServicio, &servicio.NombreServicio); err != nil {
		return domain.CrearTurnoResponse{}, domain.ErrServicioNotFound
	}

	paciente := req.Paciente
	var pacienteDB domain.PacienteRequest

	var (
		idPaciente      int
		numeroDocumento string
		apellido1       sql.NullString
		apellido2       sql.NullString
		nombre1         sql.NullString
		nombre2         sql.NullString
		sexo            int
		fechaNacimiento sql.NullString
		idTipoDocumento int
	)

	if err := r.db.QueryRow(ctx, qryPacienteByDocumento, req.Paciente.NumeroDocumento).Scan(
		&idPaciente,
		&numeroDocumento,
		&apellido1,
		&apellido2,
		&nombre1,
		&nombre2,
		&sexo,
		&fechaNacimiento,
		&idTipoDocumento,
	); err != nil {
		if !database.NoRows(err) {
			return domain.CrearTurnoResponse{}, err
		}
	} else {
		pacienteDB.IdPaciente = idPaciente
		pacienteDB.NumeroDocumento = numeroDocumento
		pacienteDB.Apellido1 = apellido1.String
		pacienteDB.Apellido2 = apellido2.String
		pacienteDB.Nombre1 = nombre1.String
		pacienteDB.Nombre2 = nombre2.String
		pacienteDB.Sexo = sexo
		pacienteDB.IdTipoDocumento = idTipoDocumento
		pacienteDB.NomTipoDocumento = ""
		pacienteDB.FechaNacimiento = fechaNacimiento.String
		paciente = pacienteDB
	}

	pacienteJSON, err := json.Marshal(paciente)
	if err != nil {
		return domain.CrearTurnoResponse{}, err
	}

	var configJSON string
	if err := r.db.QueryRow(ctx, qryConfigCantidadPorSede, req.IdSede).Scan(&configJSON); err != nil {
		if database.NoRows(err) {
			return domain.CrearTurnoResponse{}, fmt.Errorf("no hay configuración de cantidad para la sede %d", req.IdSede)
		}
		return domain.CrearTurnoResponse{}, err
	}

	if configJSON == "" {
		return domain.CrearTurnoResponse{}, fmt.Errorf("configuración de cantidad vacía para la sede %d", req.IdSede)
	}

	type cantidadConfig struct {
		CodigoModulo     json.RawMessage `json:"CodigoModulo"`
		NombreModulo     string          `json:"NombreModulo"`
		CantidadFisico   int             `json:"CantidadFisico"`
		CantidadPagina   int             `json:"CantidadPagina"`
		CantidadWhatsapp int             `json:"CantidadWhatsapp"`
	}

	var cantidadPorSede []cantidadConfig
	if err := json.Unmarshal([]byte(configJSON), &cantidadPorSede); err != nil {
		return domain.CrearTurnoResponse{}, err
	}

	var config *cantidadConfig
	for i := range cantidadPorSede {
		codigoModulo := ""
		if err := json.Unmarshal(cantidadPorSede[i].CodigoModulo, &codigoModulo); err != nil {
			var codigoNumero json.Number
			if err2 := json.Unmarshal(cantidadPorSede[i].CodigoModulo, &codigoNumero); err2 == nil {
				codigoModulo = codigoNumero.String()
			}
		}
		if codigoModulo == req.CodigoModulo {
			config = &cantidadPorSede[i]
			break
		}
	}
	if config == nil {
		return domain.CrearTurnoResponse{}, fmt.Errorf("no hay configuración de cantidad para el módulo %s", req.CodigoModulo)
	}

	var ultimoTurno int
	err = r.db.QueryRow(ctx, qryUltimoTurnoPorModuloFecha, time.Now().Format("20060102"), req.CodigoModulo).Scan(&ultimoTurno)
	if err != nil {
		if database.NoRows(err) {
			ultimoTurno = 0
		} else {
			return domain.CrearTurnoResponse{}, err
		}
	}

	var turnosHoy int
	if err := r.db.QueryRow(ctx, qryTurnosHoyPorModuloFecha, req.CodigoModulo, time.Now().Format("20060102")).Scan(&turnosHoy); err != nil {
		if database.NoRows(err) {
			turnosHoy = 0
		} else {
			return domain.CrearTurnoResponse{}, err
		}
	}

	var limite int
	switch req.Origen {
	case 1:
		limite = config.CantidadFisico
	case 2:
		limite = config.CantidadPagina
	case 3:
		limite = config.CantidadWhatsapp
	default:
		return domain.CrearTurnoResponse{}, fmt.Errorf("origen no válido: %d", req.Origen)
	}

	if turnosHoy >= limite {
		return domain.CrearTurnoResponse{}, domain.ErrTurnLimitExceeded
	}

	nuevoTurno := ultimoTurno + 1
	fechaTurno := time.Now().Format("20060102150405")

	if _, err := r.db.Exec(ctx, qryInsertTurno,
		sede.IdSede,
		sede.NomSede,
		sede.CodSede,
		compania.IdCompania,
		compania.CodigoCompania,
		compania.NombreCompania,
		tipoTurno.IdTipoTurno,
		tipoTurno.CodTipoTurno,
		tipoTurno.NomTipoTurno,
		servicio.IdServicio,
		servicio.CodigoServicio,
		servicio.NombreServicio,
		nuevoTurno,
		fechaTurno,
		0,
		0,
		string(pacienteJSON),
		paciente.IdTipoDocumento,
		req.Origen,
		req.CodigoModulo,
		req.NombreModulo,
	); err != nil {
		return domain.CrearTurnoResponse{}, err
	}

	return domain.CrearTurnoResponse{
		Accion:    domain.AccionResponse{SAccion: req.Accion},
		Sede:      sede,
		Compania:  compania,
		TipoTurno: tipoTurno,
		Servicio:  servicio,
		Paciente:  paciente,
		Turno: domain.DatosTurnoResponse{
			NumeroTurno: nuevoTurno,
			FechaTurno:  fechaTurno,
		},
	}, nil
}
