package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}
