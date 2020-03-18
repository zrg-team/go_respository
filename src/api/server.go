package api

import (
	"api/router"
	"api/utils/console"
	"auto"
	"config"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Load()
	if config.DBSYNC {
		auto.Sync()
	}
}

// Run server
func Run() {
	// fmt.Printf("\nListening [::]:%d", config.PORT)
	console.Log("\nListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
