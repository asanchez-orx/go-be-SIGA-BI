package mssql

import (
	"context"
	"encoding/json"

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
