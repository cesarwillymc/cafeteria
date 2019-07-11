package controllers

import (
	"github.com/willyrotaract/cafeteria/tree/master/productos/models"
)

type (
	// For Get - /movies
	ProductosResource struct {
		Data []models.Producto `json:"data"`
	}
	// For Post/Put - /movies
	ProductoResource struct {
		Data models.Producto `json:"data"`
	}
)
