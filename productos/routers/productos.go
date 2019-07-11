package routers

import (
	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/productos/controllers"
)

func setMovieRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/productos", controllers.GetProducto).Methods("GET")
	router.HandleFunc("/productos", controllers.CreateProducto).Methods("POST")
	router.HandleFunc("/productos/{id}", controllers.GetProductoById).Methods("GET")
	router.HandleFunc("/productos/{id}", controllers.DeleteProducto).Methods("DELETE")
	return router
}
