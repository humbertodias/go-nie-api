package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/humbertodias/go-crawler-demo/nie"
)

var provincias = nie.ScrapyProvincias()
var tramites = nie.ScrapyTramites(provincias)
var oficinas = nie.ScrapyOficinas(tramites)

func ShowApi(w http.ResponseWriter, r *http.Request) {
	endpoints := `
	<a href="/provincias">/provincias</a> <br>
	<a href="/tramites">/tramites</a> <br>
	<a href="/oficinas">/oficinas</a> <br>
	`
	fmt.Fprintln(w, endpoints)
}

func GetProvincias(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(provincias)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetTramites(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(tramites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetOficinas(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(tramites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
