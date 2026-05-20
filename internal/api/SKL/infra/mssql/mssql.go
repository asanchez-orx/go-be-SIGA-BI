package mssql

import (
	"context"

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
