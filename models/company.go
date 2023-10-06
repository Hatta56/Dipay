package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	CompanyName     string             `bson:"company_name" json:"company_name"`
	TelephoneNumber string             `bson:"telephone_number" json:"telephone_number"`
	IsActive        bool               `bson:"is_active" json:"is_active"`
	Address         string             `bson:"address" json:"address"`
}
