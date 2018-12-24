package main

import (
	"fmt"
	"net/http"

	"github.com/humbertodias/go-rest-api/handler"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", handler.ShowApi)
	http.HandleFunc("/provincias", handler.GetProvincias)
	http.HandleFunc("/tramites", handler.GetTramites)
	http.HandleFunc("/oficinas", handler.GetOficinas)

	fmt.Printf("Listening at http://0.0.0.0:%d\n", PORT)
	addr := fmt.Sprintf(":%d", PORT)
	http.ListenAndServe(addr, nil)
}
