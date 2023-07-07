package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID      primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	CarName string             `json:"carName" bson:"carName"`
	Company string             `json:"company" bson:"company"`
	MadeAt  time.Time          `json:"madeAt" bson:"madeAt,omitempty"`
	SoldAt  time.Time          `json:"soldAt" bson:"soldAt,omitempty"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
