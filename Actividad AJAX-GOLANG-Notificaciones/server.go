package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type movies struct {
	Titulo       string  `json: titulo`
	Director     string  `json: director`
	Duracion     int     `json: duracion`
	Calificacion float64 `json: calificacion`
}

var pelis = []movies{
	{
		Titulo:       "Joker",
		Director:     "Todd Phillips",
		Duracion:     121,
		Calificacion: 8.4,
	},
	{
		Titulo:       "Cuties",
		Director:     "Maïmouna Doucouré",
		Duracion:     95,
		Calificacion: 10,
	},
	{
		Titulo:       "Fiebre de Amor",
		Director:     "Rene Cardona Jr.",
		Duracion:     97,
		Calificacion: 9,
	},
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "actividad.html")
}

func reto(w http.ResponseWriter, r *http.Request) {

	arreglo := [6]int{2, 3, 5, 7, 11, 13}

	switch r.Method {
	case "GET":
		rand.Seed(time.Now().UnixNano())
		position := rand.Intn(3)
		json.NewEncoder(w).Encode(arreglo[position])

	case "POST":
		r.ParseForm()
		position, _ := strconv.ParseInt(r.Form.Get("y"), 10, 64)
		dato, _ := strconv.Atoi(r.Form.Get("n"))

		arreglo[position] = dato
		fmt.Println(arreglo)
	}

}

func mostrarJSON(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(pelis)

}

func insertarMovie(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pTitulo := r.Form.Get("titulo")
	pDirector := r.Form.Get("director")
	pDuracion, _ := strconv.Atoi(r.Form.Get("duracion"))
	pCalificacion, _ := strconv.Atoi(r.Form.Get("calificacion"))

	peli := movies{
		Titulo:       pTitulo,
		Director:     pDirector,
		Duracion:     pDuracion,
		Calificacion: float64(pCalificacion),
	}

	pelis = append(pelis, peli)

}

func main() {

	http.HandleFunc("/reto", reto)
	http.HandleFunc("/actividad2", mostrarJSON)
	http.HandleFunc("/insertar", insertarMovie)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8081", nil)
	if err != nil {
		return
	}
}
