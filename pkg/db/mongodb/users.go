package mongodb

import (
	"context"

	"github.com/andreylm/bmlabs/pkg/entities"
	"github.com/andreylm/bmlabs/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userCollection string = "user"

type userRepository struct {
	db *mongo.Database
}

func (u *userRepository) Create(ctx context.Context, user *entities.User) (err error) {
	collection := u.db.Collection(userCollection)
	if _, err = collection.InsertOne(ctx, user); err != nil {
		logger.Get().Infof("Error inserting user: %v", user)
		return
	}

	logger.Get().Info("Inserted user into collection")
	return
}

func (u *userRepository) Search(ctx context.Context, limit, offset *int64) (users []*entities.User, err error) {
	collection := u.db.Collection(userCollection)
	var opts []*options.FindOptions

	opt := options.FindOptions{
		Limit: limit,
		Skip:  offset,
	}
	opts = append(opts, &opt)

	cursor, err := collection.Find(ctx, struct{}{}, opts...)
	if err != nil {
		logger.Get().Error(err)
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		logger.Get().Error(err)
	}

	return
}
