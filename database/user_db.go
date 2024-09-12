package database

import (
	"context"
	"log"

	"github.com/renpereiradx/marvel-api/model"
)

func (pgrp *PostgresRepo) InsertUser(ctx context.Context, user *model.User) error {
	_, err := pgrp.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password)
	return err
}
func (pgrp *PostgresRepo) GetUser(ctx context.Context, email string) (*model.User, error) {
	rows, err := pgrp.db.QueryContext(ctx, "SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println("error closing rows", err)
		}
	}()

	var user model.User
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.Created_At, &user.Changed_At); err != nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (pgrp *PostgresRepo) UpdateUser(ctx context.Context, user *model.User) error {
	_, err := pgrp.db.ExecContext(ctx, "UPDATE users id = $1, email = $2, changed_at = NOW()", user.ID, user.Email)
	return err
}

func (pgrp *PostgresRepo) DeleteUser(ctx context.Context, email string) error {
	_, err := pgrp.db.ExecContext(ctx, "DELETE FROM users WHERE email = $1", email)
	return err
}
