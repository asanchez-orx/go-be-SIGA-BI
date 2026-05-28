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
