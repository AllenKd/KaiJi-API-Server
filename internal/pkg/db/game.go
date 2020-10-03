package db

import (
	"KaiJi-API-Server/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *client) GetGames(filter bson.M, option *options.FindOptions) ([]collection.SportsData, error) {
	log.Debug("query games from db: ", filter)

	cursor, err := c.SportsData.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get document: ", err.Error())
		return nil, err
	}
	var documents []collection.SportsData
	if err := cursor.All(nil, &documents); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return documents, nil
}

func (c *client) CountGames(filter bson.M) int64 {
	count, err := c.SportsData.CountDocuments(nil, filter)
	if err != nil {
		log.Error("fail to count documents: ", err.Error())
		return -1
	}
	return count
}
