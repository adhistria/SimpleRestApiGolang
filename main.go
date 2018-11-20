package main

import (
	"log"
	"net/http"
	"omahihrom/route"
)


func main() {

	log.Fatal(http.ListenAndServe(":8000", route.Router))
}
