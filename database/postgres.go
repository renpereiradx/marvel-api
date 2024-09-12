package database

import (
	"database/sql"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(url string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{db: db}, nil
}

// pgrp = Postgres Repository
func (pgrp *PostgresRepo) Close() error {
	return pgrp.db.Close()
}
