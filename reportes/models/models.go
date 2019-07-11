package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Reporte struct {
		Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId     string        `json:"userid"`
		VentaID    []string        `json:"ventas"`
		Total     int      `json:"total"`
	}
)
