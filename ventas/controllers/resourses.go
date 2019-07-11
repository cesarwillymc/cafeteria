package controllers

import (
	"github.com/willyrotaract/cafeteria/ventas/models"
)

type (
	// For Get - /showtimes
	VentasResource struct {
		Data []models.Venta `json:"data"`
	}
	// For Post/Put - /showtimes
	VentaResource struct {
		Data models.Venta `json:"data"`
	}
)
