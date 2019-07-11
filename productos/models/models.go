package models

import (

	"gopkg.in/mgo.v2/bson"
)

type (
	Producto struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title     string        `json:"title"`
		Marca     string         `json:"Marca"`
		Precio    float32       `json:"Precio"`
		Tipo      string	`json:"Tipo"`
	}
)
