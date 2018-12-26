package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/humbertodias/go-nie-crawler/nie"
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
	ToJson(provincias, w)
}

func GetTramites(w http.ResponseWriter, r *http.Request) {
	ToJson(tramites, w)
}

func GetOficinas(w http.ResponseWriter, r *http.Request) {
	ToJson(oficinas, w)
}

func ToJson(arr interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(arr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}