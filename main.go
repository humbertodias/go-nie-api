package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/humbertodias/go-nie-api/handler"
)

func main() {
	PORT := "8080"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}
	go handler.Init()
	http.HandleFunc("/", handler.ShowApi)
	http.HandleFunc("/provincias", handler.GetProvincias)
	http.HandleFunc("/tramites", handler.GetTramites)
	http.HandleFunc("/oficinas", handler.GetOficinas)

	fmt.Printf("Listening at http://0.0.0.0:%s\n", PORT)
	addr := fmt.Sprintf(":%s", PORT)
	http.ListenAndServe(addr, nil)
}
