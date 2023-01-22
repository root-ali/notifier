package main

import (
	"log"
	"notifier/pkg/http/rest"
	"os"
)

func main() {
	// Creating grafana and alertmanager Service

	// start http server
	router := rest.Handler()
	http_port := "9090"
	if value, ok := os.LookupEnv("HTTP_PORT"); ok {
		http_port = value
	}
	log.Fatal(router.Run(":" + http_port))

}
