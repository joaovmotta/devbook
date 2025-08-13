package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.ConfigureEnvironment()

	router := router.Generate()

	fmt.Println("Api running on port", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
