package db

import (
	"context"

	"github.com/andreylm/bmlabs/pkg/entities"
)

type Storer interface {
	User() UserRepository
}

type UserRepository interface {
	Create(context.Context, entities.User) error
	Search(context.Context) ([]entities.User, error)
}
