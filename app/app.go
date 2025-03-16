package main

import (
	"back/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.Routing()
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
