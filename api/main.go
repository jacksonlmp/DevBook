package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running API")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))
}
