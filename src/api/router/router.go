package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

// New app route
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
