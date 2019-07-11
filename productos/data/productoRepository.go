package data

import (
	"time"

	"github.com/willyrotaract/cafeteria/tree/master/productos/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductoRepository struct {
	C *mgo.Collection
}

func (r *ProductoRepository) GetAll() []models.Producto {
	var Productos []models.Producto
	iter := r.C.Find(nil).Iter()
	result := models.Producto{}
	for iter.Next(&result) {
		productos = append(productos, result)
	}
	return productos
}

func (r *ProductoRepository) Create(producto *models.Producto) error {
	obj_id := bson.NewObjectId()
	producto.Id = obj_id
	producto.CreatedOn = time.Now()
	err := r.C.Insert(&producto)
	return err
}

func (r *ProductoRepository) GetById(id string) (producto models.Producto, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&producto)
	return
}

func (r *ProductoRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
