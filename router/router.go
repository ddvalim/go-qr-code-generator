package router

import (
	"github.com/ddvalim/go-qr-code-generator/core/ports"
	"github.com/ddvalim/go-qr-code-generator/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()

	var r []ports.Route

	r = append(r, routes.QRCode...)

	for _, route := range r {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return router
}
