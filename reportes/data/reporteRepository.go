package data

import (
	"github.com/willyrotaract/cafeteria/tree/master/reportes/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReporteRepository struct {
	C *mgo.Collection
}

func (r *ReporteRepository) Create(Reporte *models.Reporte) error {
	obj_id := bson.NewObjectId()
	Reporte.Id = obj_id
	err := r.C.Insert(&Reporte)
	return err
}

func (r *ReporteRepository) GetAll() []models.Reporte {
	var Reportes []models.Reporte
	iter := r.C.Find(nil).Iter()
	result := models.Reporte{}
	for iter.Next(&result) {
		Reportes = append(Reportes, result)
	}
	return Reportes
}

func (r *ReporteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
