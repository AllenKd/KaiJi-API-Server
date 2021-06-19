package db

import (
	"KaiJi-Admin/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c client) GetUser(userName string) (userData collection.User, err error) {
	log.Debug("get user: ", userName)

	filter := bson.M{
		"name": userName,
	}
	err = c.User.FindOne(nil, filter).Decode(&userData)
	return
}
