package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       *primitive.ObjectID `bson:"_id,omitempty"`
	Name     string              `bson:"name" binding:"required"`
	Password string              `bson:"password" binding:"required"` // hashed password
	//Permission string
}
