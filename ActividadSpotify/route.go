package main

import (
	"encoding/json"
	"net/http"

	"github.com/Gussiny/apiRest/repo"
)

var (
	repository repo.UserRepository = repo.NewUserRepository()
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func get(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)
	res.Header().Set("Content-type", "application/json")
	users, errRepo := repository.FindAll()
	if errRepo != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error getting the users from Firestore"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(users)
}

func post(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "Archivo publicado"}`))
}

func patch(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "Archivo patcheado"}`))
}

func put(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "Archivo puteado" }`))
}

func delete(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"message": "Archivo deleteado"}`))

}
