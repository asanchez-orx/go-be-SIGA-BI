package mssql

import (
	"context"
	"encoding/json"
	"strings"
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
		rows, err = r.db.Query(ctx, qryTipoServicio, req.CodigoModulo, req.IdSede)
	} else {
		rows, err = r.db.Query(ctx, qryTipoServicioPorCompania, req.IdCompania, req.CodigoModulo, req.IdSede)
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
		if err := rows.Scan(&item.IdSede, &item.CodSede, &item.NomSede); err != nil {
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

func (r *CreacionTurnosRepo) CrearTurno(ctx context.Context, req domain.CrearTurnoRequest) (domain.CrearTurnoResponse, error) {
	var resp domain.CrearTurnoResponse
	resp.Accion.SAccion = req.Accion

	// 1. Single roundtrip to validate Sede, Compania, TipoTurno, Servicio and get quantities json configurations
	var nomSede, codSede, codCompania, nomCompania, codTipoTurno, nomTipoTurno, codServicio, nomServicio, jsonCantidad *string

	err := r.db.QueryRow(ctx, qryValidarYConfigurarTurno, req.IdSede, req.IdCompania, req.IdTipoTurno, req.IdServicio).Scan(
		&nomSede, &codSede,
		&codCompania, &nomCompania,
		&codTipoTurno, &nomTipoTurno,
		&codServicio, &nomServicio,
		&jsonCantidad,
	)
	if err != nil {
		return resp, err
	}

	// Validate scanned results to make sure all entities exist in the database
	if nomSede == nil || codSede == nil {
		return resp, domain.ErrSedeNotFound
	}
	resp.Sede.IdSede = req.IdSede
	resp.Sede.NomSede = *nomSede
	resp.Sede.CodSede = *codSede

	if codCompania == nil || nomCompania == nil {
		return resp, domain.ErrCompaniaNotFound
	}
	resp.Compania.IdCompania = req.IdCompania
	resp.Compania.CodigoCompania = *codCompania
	resp.Compania.NombreCompania = *nomCompania

	if codTipoTurno == nil || nomTipoTurno == nil {
		return resp, domain.ErrTipoTurnoNotFound
	}
	resp.TipoTurno.IdTipoTurno = req.IdTipoTurno
	resp.TipoTurno.CodTipoTurno = *codTipoTurno
	resp.TipoTurno.NomTipoTurno = *nomTipoTurno

	if codServicio == nil || nomServicio == nil {
		return resp, domain.ErrServicioNotFound
	}
	resp.Servicio.IdServicio = req.IdServicio
	resp.Servicio.CodigoServicio = *codServicio
	resp.Servicio.NombreServicio = *nomServicio

	// 2. Fetch/Validate Paciente
	err = r.db.QueryRow(ctx, qryValidarPacienteTurno, req.Paciente.NumeroDocumento).Scan(
		&resp.Paciente.IdPaciente, &resp.Paciente.NumeroDocumento, &resp.Paciente.Apellido1, &resp.Paciente.Apellido2,
		&resp.Paciente.Nombre1, &resp.Paciente.Nombre2, &resp.Paciente.Sexo, &resp.Paciente.FechaNacimiento,
		&resp.Paciente.IdTipoDocumento, &resp.Paciente.NomTipoDocumento)
	
	if err != nil {
		if database.NoRows(err) {
			// If not found, use the one from request
			resp.Paciente = req.Paciente
		} else {
			return resp, err
		}
	}

	// 3. Parse limits from Sede config
	var cantidades []domain.CantidadModulo
	if jsonCantidad != nil && *jsonCantidad != "" {
		_ = json.Unmarshal([]byte(*jsonCantidad), &cantidades)
	}

	var cantFisico, cantPagina, cantWhatsapp int
	for _, c := range cantidades {
		if c.CodigoModulo == req.CodigoModulo {
			cantFisico = c.CantidadFisico
			cantPagina = c.CantidadPagina
			cantWhatsapp = c.CantidadWhatsapp
			break
		}
	}

	// Determine the limit depending on the Origin
	var limite int
	switch req.Origen {
	case 1:
		limite = cantFisico
	case 2:
		limite = cantPagina
	case 3:
		limite = cantWhatsapp
	default:
		limite = 0
	}

	// Convert paciente to JSON to save it
	jsonPacienteBytes, _ := json.Marshal(resp.Paciente)
	jsonPaciente := string(jsonPacienteBytes)

	// Format current local time in WinDev format: YYYYMMDDHHmmss
	horaActual := time.Now().Format("20060102150405")

	// 4. Concurrency-safe atomic insert and fetch generated fields (all in exactly one DB call!)
	err = r.db.QueryRow(ctx, qryInsertarTurnoConsolidado,
		resp.Sede.IdSede, resp.Sede.NomSede, resp.Sede.CodSede,
		resp.Servicio.IdServicio, resp.Servicio.CodigoServicio, resp.Servicio.NombreServicio,
		resp.TipoTurno.IdTipoTurno, resp.TipoTurno.CodTipoTurno, resp.TipoTurno.NomTipoTurno,
		resp.Compania.IdCompania, resp.Compania.CodigoCompania, resp.Compania.NombreCompania,
		horaActual, jsonPaciente, resp.Paciente.IdTipoDocumento,
		req.Origen, req.CodigoModulo, req.NombreModulo, limite).Scan(&resp.Turno.NumeroTurno, &resp.Turno.FechaTurno)

	if err != nil {
		if strings.Contains(err.Error(), "LIMIT_EXCEEDED") {
			return resp, domain.ErrTurnLimitExceeded
		}
		return resp, err
	}

	return resp, nil
}
