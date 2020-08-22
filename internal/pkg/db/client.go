package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type client struct {

}

var (
	once sync.Once
	instance *mongo.Client
)

func New() *mongo.Client {
	once.Do(func() {
		instance = mongo.Connect(nil, options.Client().ApplyURI() )
	})
}