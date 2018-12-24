package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/humbertodias/go-crawler-demo/nie"
)

var oficinas = nie.ScrapyOficinas()
var tramites = nie.ScrapyTramites(oficinas)

func ShowApi(w http.ResponseWriter, r *http.Request) {
	endpoints := `
	<a href="/oficinas">/oficinas</a> <br>
	<a href="/tramites">/tramites</a> <br>
	`
	fmt.Fprintln(w, endpoints)
}

func GetOficinas(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(oficinas)
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
