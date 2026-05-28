package mssql

import (
	"context"

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
