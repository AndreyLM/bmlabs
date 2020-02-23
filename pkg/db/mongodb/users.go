package mongodb

import (
	"context"

	"github.com/andreylm/bmlabs/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func (u *userRepository) Create(context.Context, entities.User) error {
	return nil
}

func (u *userRepository) Search(context.Context) ([]entities.User, error) {
	return nil, nil
}
