package db

import (
	"context"

	"github.com/andreylm/bmlabs/pkg/entities"
)

// Storer - container for all repositories
type Storer interface {
	User() UserRepository
}

// UserRepository - interaction with db's user table or collection
type UserRepository interface {
	Create(context.Context, *entities.User) error
	Search(ctx context.Context, limit, offset *int64) ([]*entities.User, error)
}
