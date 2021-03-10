package main

import (
	"encoding/json"
	"net/http"
)

type movies struct {
	Titulo   string `json: "titulo"`
	Director string `json: "director"`
	Cantidad int    `json: "cantidad"`
	Genero   string `json: "genero"`
}

type book struct {
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Oferta    bool   `json:"oferta"`
	Cantidad  int    `json:"cantidad"`
	Editorial string `json:"editorial"`
}

var pelis = []movies{
	{
		Titulo:   "Joker",
		Director: "Todd Phillips",
		Cantidad: 20,
		Genero:   "Suspenso",
	},
	{
		Titulo:   "Cuties",
		Director: "Maïmouna Doucouré",
		Cantidad: 30,
		Genero:   "Drama",
	},
	{
		Titulo:   "Fiebre de Amor",
		Director: "Rene Cardona Jr.",
		Cantidad: 25,
		Genero:   "Musical",
	},
}

var libros = []book{
	{
		Titulo:    "Fahrenheit 451",
		Autor:     "Ray Bradbury",
		Oferta:    true,
		Cantidad:  14,
		Editorial: "Ballantine Books",
	},
	{
		Titulo:    "El fin de la infancia",
		Autor:     "Arthur C. Clarke",
		Oferta:    false,
		Cantidad:  21,
		Editorial: "Ballantine Books",
	},
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "RetoVue.html")
}

func mostrarPelis(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(pelis)
	}

}

func mostrarLibros(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(libros)
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	http.HandleFunc("/movies", mostrarPelis)
	http.HandleFunc("/books", mostrarLibros)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8081", nil)
	if err != nil {
		return
	}
}
