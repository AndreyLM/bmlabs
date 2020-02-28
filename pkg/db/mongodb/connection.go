package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/andreylm/bmlabs/pkg/db"
	"github.com/andreylm/bmlabs/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoStorage struct {
	user db.UserRepository
}

// NewStorage - creates mongo db storage
func NewStorage(host string, port int, dbName, user, password string) db.Storer {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Get().Fatal("Cannot initialize database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.Get().Fatal("Cannot initialize database context")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Get().Fatal("Cannot ping database")
	}

	logger.Get().Info("Connected To MongoDB")

	db := client.Database(dbName)
	mStorer := &mongoStorage{
		user: &userRepository{db},
	}

	return mStorer
}

func (s *mongoStorage) User() db.UserRepository {
	return s.user
}
