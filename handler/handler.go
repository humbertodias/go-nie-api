package handler

import (
	"encoding/json"
	"fmt"
	"github.com/humbertodias/go-nie-api/fake"
	"net/http"

	. "github.com/humbertodias/go-nie-crawler/model"
	"github.com/humbertodias/go-nie-crawler/nie"
)

var provincias []Provincia
var tramites []Tramite
var oficinas []Oficina

func Init() {
	provincias = nie.ScrapyProvincias()
	tramites = nie.ScrapyTramites(provincias)
	oficinas = nie.ScrapyOficinas(tramites)
}

func ShowApi(w http.ResponseWriter, r *http.Request) {
	endpoints := `
	<a href="/provincias">/provincias</a> <br>
	<a href="/tramites">/tramites</a> <br>
	<a href="/oficinas">/oficinas</a> <br>
	<a href="/nif">/nif</a> <br>
	<a href="/nie">/nie</a> <br>
	<a href="/dni">/dni</a> <br>
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

func RandomNif(w http.ResponseWriter, r *http.Request) {
	ToJson(fake.NifRandom(), w)
}

func RandomNie(w http.ResponseWriter, r *http.Request) {
	ToJson(fake.NieRandom(), w)
}

func RandomDni(w http.ResponseWriter, r *http.Request) {
	ToJson(fake.DniRandom(), w)
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
