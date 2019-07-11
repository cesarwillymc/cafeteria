package routers

import (
	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/reportes/controllers"
)

func SetReportesRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/reportes", controllers.GetReportes).Methods("GET")
	router.HandleFunc("/reportes", controllers.CreateReportes).Methods("POST")
	return router
}
