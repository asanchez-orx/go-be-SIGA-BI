package postgres

import (
	"context"
	"strconv"

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

// Métodos vacíos (stubs) requeridos por la interfaz CreacionTurnosRepository
func (r *CreacionTurnosRepo) BuscarCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) (domain.CreacionTurnosesResponse, error) {
	return domain.CreacionTurnosesResponse{}, nil
}

func (r *CreacionTurnosRepo) CrearCreacionTurnos(ctx context.Context, req domain.CreacionTurnosRequest) error {
	return nil
}

func (r *CreacionTurnosRepo) BuscarCompania(ctx context.Context, req domain.CompaniaRequest) (domain.CompaniasResponse, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) VerificarConfigCompanias(ctx context.Context) (domain.ConfigCompaniasResponse, error) {
	var res domain.ConfigCompaniasResponse
	var manejaEmpresasStr string

	err := r.db.QueryRow(ctx, qryVerificarConfigCompanias).Scan(&manejaEmpresasStr)
	if err != nil {
		if database.NoRows(err) {
			return domain.ConfigCompaniasResponse{ManejaEmpresas: 0}, nil
		}
		return domain.ConfigCompaniasResponse{}, err
	}

	val, _ := strconv.Atoi(manejaEmpresasStr)
	res.ManejaEmpresas = val

	return res, nil
}

func (r *CreacionTurnosRepo) BuscarTipoServicio(ctx context.Context, req domain.TipoServicioRequest) (domain.TipoServiciosResponse, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) BuscarModulo(ctx context.Context, req domain.ModuloRequest) (domain.ModulosResponse, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) BuscarSedes(ctx context.Context) (domain.SedesResponse, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) BuscarTipoTurno(ctx context.Context, req domain.TipoTurnoRequest) (domain.TipoTurnosResponse, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) ConfirmarConfigSedes(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *CreacionTurnosRepo) CrearTurno(ctx context.Context, req domain.CrearTurnoRequest) (domain.CrearTurnoResponse, error) {
	return domain.CrearTurnoResponse{}, nil
}

func (r *CreacionTurnosRepo) CargarConfigLIS(ctx context.Context) (domain.ConfigLISResponse, error) {
	return domain.ConfigLISResponse{}, nil
}
