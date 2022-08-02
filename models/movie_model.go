package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	Fullplot string             `json:"fullplot,omitempty" validate:"required"`
	Year     int                `json:"year,omitempty" validate:"required"`
}
