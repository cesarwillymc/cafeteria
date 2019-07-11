package controllers

import (
	"github.com/willyrotaract/cafeteria/tree/master/reportes/models"
)

type (
	// For Get - /bookings
	ReportesResource struct {
		Data []models.Reporte `json:"data"`
	}
	// For Post/Put - /bookings
	ReporteResource struct {
		Data models.Reporte `json:"data"`
	}
)
