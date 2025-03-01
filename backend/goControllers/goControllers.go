package goControllers

import (
	"encoding/json"
	"goCrud/goHandlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(router *mux.Router) {
	router.HandleFunc("/LoadData", LoadData).Methods("GET")
	router.HandleFunc("/SaveData", SaveData).Methods("PUT")
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/jsonO")
	users, err := goHandlers.LoadDataHandler("../database/data.json")
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding users to JSON", http.StatusInternalServerError)
		return
	}
}

func SaveData(w http.ResponseWriter, r *http.Request) {
	var newUser goHandlers.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	err = goHandlers.SaveDataHandler("../database/data.json", newUser)
	if err != nil {
		http.Error(w, "Isue creating user data", http.StatusBadRequest)
		return
	}

	// Succesfully Created User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
