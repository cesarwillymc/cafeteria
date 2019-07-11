package routers

import (
	"github.com/gorilla/mux"
	"github.com/willyrotaract/cafeteria/tree/master/reportes/controllers"
)

func SetReportesRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/reportes", controllers.GetReportes).Methods("GET")
	router.HandleFunc("/reportes", controllers.CreateReporte).Methods("POST")
	return router
}
