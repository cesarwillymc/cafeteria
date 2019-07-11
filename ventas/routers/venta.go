package routers

import (
	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/ventas/controllers"
)

func SetVentaRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/ventas", controllers.GetVenta).Methods("GET")
	router.HandleFunc("/ventas", controllers.CreateVenta).Methods("POST")
	router.HandleFunc("/ventas/{date}", controllers.GetVentaByDate).Methods("GET")
	router.HandleFunc("/ventas/{id}", controllers.DeleteVenta).Methods("DELETE")
	return router
}
