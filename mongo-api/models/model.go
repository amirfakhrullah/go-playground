package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Movie   string             `bson:"movie,omitempty" json:"movie"`
	Watched bool               `bson:"watched,omitempty" json:"watched"`
}
