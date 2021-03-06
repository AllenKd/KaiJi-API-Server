package db

import (
	"KaiJi-Admin/internal/pkg/util"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) GetUser(userName string) (userData structs.User, err error) {
	log.Debug("get user: ", userName)

	filter := bson.M{
		"name": userName,
	}
	err = c.User.FindOne(nil, filter).Decode(&userData)
	return
}

func (c client) CreateUser(username, password string) (userData structs.User, err error) {
	userData = structs.User{
		Name:     username,
		Password: util.HashedPassword(password),
	}
	res, dbErr := c.User.InsertOne(nil, userData)
	if dbErr != nil {
		log.Error("fail to insert document: ", dbErr.Error())
		err = dbErr
		return
	}
	id := res.InsertedID.(primitive.ObjectID)
	userData.Id = &id
	return
}
