package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        *string            `json:"name" validate:"required,min=2,max=60"`
	RoleName    *string            `json:"roleName" validate:"required,min=60`
	Description *string            `json:"description" validate:"min=2,max=100"`
}
