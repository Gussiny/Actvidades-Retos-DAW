package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//$env:GOOGLE_APPLICATION_CREDENTIALS="C:\Users\hello\OneDrive - Instituto Tecnologico y de Estudios Superiores de Monterrey\ITESM\8VO\DesarrolloWeb\Servidor\Retos\ActividadSpotify\spotifypirata-b35dd-firebase-adminsdk-3z0ak-fa43f8cb2c.json"

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(res, "Up and running")
	})

	router.HandleFunc("/users", get).Methods("GET")
	router.HandleFunc("/users", post).Methods("POST")
	router.HandleFunc("/users", patch).Methods("PATCH")
	router.HandleFunc("/users", put).Methods("PUT")
	router.HandleFunc("/users", delete).Methods("DELETE")

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
