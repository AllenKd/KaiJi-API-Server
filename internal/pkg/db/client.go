package db

import (
	"dataProvider/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	once sync.Once
	instance *mongo.Client
)

func New() *mongo.Client {
	once.Do(func() {
		client, err := mongo.Connect(nil, options.Client().ApplyURI(configs.New().Db.ConnectionString))
		if err != nil {
			panic(err)
		}
		instance = client
		log.Debug("mongo client initialized")
	})
	return instance
}