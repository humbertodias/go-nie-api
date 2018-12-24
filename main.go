package main

import (
	"fmt"
	"net/http"

	"github.com/humbertodias/go-rest-api/handler"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", handler.ShowApi)
	http.HandleFunc("/oficinas", handler.GetOficinas)
	http.HandleFunc("/tramites", handler.GetTramites)

	fmt.Printf("Listening at http://0.0.0.0:%d\n", PORT)
	addr := fmt.Sprintf(":%d", PORT)
	http.ListenAndServe(addr, nil)
}
