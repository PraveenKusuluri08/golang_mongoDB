package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PraveenKusuluri08/routes"
)

func main() {
	fmt.Println("USER->MONGODB🍀")
	r := routes.Router()
	fmt.Println("App is listining")
	log.Fatal(http.ListenAndServe(":5000", r))
}
