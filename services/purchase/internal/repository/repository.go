package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pgPool   *pgxpool.Pool
	txMarker string
}

func NewRepository(pgPool *pgxpool.Pool, txMarker string) *Repository {
	return &Repository{
		pgPool:   pgPool,
		txMarker: txMarker,
	}
}


