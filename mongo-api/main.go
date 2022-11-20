package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/amirfakhrullah/go-playground/tree/main/mongo-api/routers"
)

func main() {
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is listening on PORT 8080")
}
