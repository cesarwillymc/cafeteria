package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/ventas/common"
	"github.com/willyrotaract/cafeteria/ventas/data"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/showtimes"
// Returns all Showtime documents
func GetVentas(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ventas")
	repo := &data.VentaRepository{c}
	// Get all showtimes form repository
	ventas := repo.GetAll()
	j, err := json.Marshal(VentasResource{Data: ventas})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/showtimes"
// Create a new Showtime document
func CreateVenta(w http.ResponseWriter, r *http.Request) {
	var dataResource VentaResource
	// Decode the incoming ShowTime json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid ShowTime data", 500)
		return
	}
	venta := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ventas")
	// Create ShowTime
	repo := &data.VentaRepository{c}
	repo.Create(venta)
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetVentaByDate(w http.ResponseWriter, r *http.Request) {
	// Get date from incoming url
	vars := mux.Vars(r)
	date := vars["date"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ventas")
	repo := &data.VentaRepository{c}

	// Get showtime by date
	venta, err := repo.GetByDate(date)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Create data for the response
	j, err := json.Marshal(VentaResource{Data: venta})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/showtimes/{id}"
// Delete a Showtime document by id
func DeleteVenta(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("ventas")

	// Remove showtime by id
	repo := &data.VentaRepository{c}
	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error ahs occurred", 500)
			return
		}
	}

	// Send response back
	w.WriteHeader(http.StatusNoContent)
}
