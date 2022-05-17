package main

import (
	"fmt"
	"log"
	"net/http"
	"stocks/router"
)

func main() {
	fmt.Println("Using PostgreSQL wiht GO..")

	r := router.Router()
	fmt.Println("Starting server on the port 8081..")
	log.Fatal(http.ListenAndServe(":5432", r))
}
