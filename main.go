package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type application struct {
	Port   string
	Domain string
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payloud = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "True",
		Message: "Golang RestApi",
	}

	jsonByte, err := json.Marshal(payloud)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonByte)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func main() {
	var app application
	app.Port = ":8080"

	mux := chi.NewRouter()
	mux.Get("/", app.Home)

	fmt.Println("Starting application on port ", app.Port)
	err := http.ListenAndServe(app.Port, mux)
	if err != nil {
		panic(err)
	}

}
