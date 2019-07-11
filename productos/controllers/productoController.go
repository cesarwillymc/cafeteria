package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/productos/common"
	"github.com/willyrotaract/cafeteria/productos/data"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/movies"
// Returns all Movie documents
func GetProductos(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("productos")
	repo := &data.ProductoRepository{c}
	productos := repo.GetAll()
	j, err := json.Marshal(ProductosResource{Data: productos})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/movies"
// Insert a new Movie document
func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var dataResourse ProductoResource
	// Decode the incoming Movie json
	err := json.NewDecoder(r.Body).Decode(&dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Movie data", 500)
		return
	}
	producto := &dataResourse.Data

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("productos")
	// Insert a movie document
	repo := &data.ProductoRepository{c}
	repo.Create(producto)
	j, err := json.Marshal(dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Get - "/movies/{id}"
// Get movie by id
func GetProductoById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("movies")
	repo := &data.ProductoRepository{c}

	// Get movie by id
	producto, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(ProductoResource{Data: producto})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/movies/{id}"
// Delete movie by id
func DeleteProducto(rw http.ResponseWriter, req *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(req)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("productos")
	repo := &data.ProductoRepository{c}

	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(rw, err, "An unexpected error has occurred", 500)
			return
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}
