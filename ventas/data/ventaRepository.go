package data

import (
	"time"

	"github.com/willyrotaract/cafeteria/ventas/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VentaRepository struct {
	C *mgo.Collection
}

func (r *VentaRepository) Create(Venta *models.Venta) error {
	obj_id := bson.NewObjectId()
	Venta.Id = obj_id
	Venta.CreatedOn = time.Now()
	err := r.C.Insert(&Venta)
	return err
}

func (r *VentaRepository) GetAll() []models.Venta {
	var Ventas []models.Venta
	iter := r.C.Find(nil).Iter()
	result := models.Venta{}
	for iter.Next(&result) {
		Ventas = append(Ventas, result)
	}
	return Ventas
}

func (r *VentaRepository) GetByDate(date string) (venta models.Venta, err error) {
	err = r.C.Find(bson.M{"date": date}).One(&venta)
	return
}

func (r *VentaRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
