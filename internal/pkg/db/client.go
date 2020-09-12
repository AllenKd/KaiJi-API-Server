package db

import (
	"apiServer/internal/pkg/configs"
	"apiServer/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type client struct {
	*mongo.Client
	SportsData *mongo.Collection
}

var (
	once sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(configs.New().Mongo.ConnectionString))
		if err != nil {
			panic(err)
		}
		db := c.Database(configs.New().Mongo.Db)
		instance = &client{
			c,
			db.Collection("sports_data"),
		}
		if err := instance.Ping(nil, nil); err != nil {
			panic(err)
		}
		log.Debug("mongo client initialized")
	})
	return instance
}

func (c *client) GetGames(filter bson.M) ([]collection.SportsData, error){
	log.Debug("query games from db: ", filter)
	var documents []collection.SportsData
	cursor, err := c.SportsData.Find(nil, filter)
	if err != nil {
		log.Error("fail to get document: ", err.Error())
		return nil, err
	}
	if err := cursor.All(nil, &documents); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return documents, nil
}