package routes

import (
	"api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		URI:          "/version",
		Method:       http.MethodGet,
		Handler:      controllers.Version,
		AuthRequired: false,
	},
}
