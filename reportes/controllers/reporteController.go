package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/willyrotaract/cafeteria/tree/master/reportes/common"
	"github.com/willyrotaract/cafeteria/tree/master/reportes/data"
)

// Handler for HTTP Post - "/bookins"
// Create a new Booking document
func CreateReporte(w http.ResponseWriter, r *http.Request) {
	var dataResource ReporteResource
	// Decode the incoming Booking json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid reporte data", 500)
		return
	}
	Reporte := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("Reportes")
	// Create Booking
	repo := &data.ReporteRepository{c}
	repo.Create(Reporte)
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

func GetReportes(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("Reportes")
	repo := &data.ReporteRepository{c}
	// Get all bookings
	reportes := repo.GetAll()
	// Create response data
	j, err := json.Marshal(ReportesResource{Data: reportes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
