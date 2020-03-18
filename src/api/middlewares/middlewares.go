package middlewares

import (
	"api/auth"
	"api/utils/console"
	"encoding/json"
	"net/http"
	"time"
)

// SetMiddlewareLogger displays a info message of the API
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		console.Highlight("[%s] %s %s%s %s", currentTime.Format("2006-01-02 15:04:05"), r.Method, r.Host, r.RequestURI, r.Proto)

		next(w, r)
	}
}

// SetMiddlewareJSON set the application Content-Type
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication authorize an access
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			res := make(map[string]interface{})
			res["message"] = "Unauthorized."
			w.WriteHeader(http.StatusForbidden)
			jData, _ := json.Marshal(res)
			w.Write(jData)

			return
		}
		next(w, r)
	}
}
