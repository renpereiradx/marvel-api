package repository

import (
	"context"

	"github.com/renpereiradx/marvel-api/model"
)

type Repository interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, email string) error

	Close() error
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func InsertUser(ctx context.Context, user *model.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUser(ctx context.Context, email string) (*model.User, error) {
	return implementation.GetUser(ctx, email)
}

func UpdateUser(ctx context.Context, user *model.User) error {
	return implementation.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, email string) error {
	return implementation.DeleteUser(ctx, email)
}

func Close() error {
	return implementation.Close()
}
