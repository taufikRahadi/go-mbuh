package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quiz-1/infrastructure/application/common/middlewares"
	"quiz-1/infrastructure/config"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

func main() {
	server := mux.NewRouter()
	dbConn, err := config.DbConn()

	if err != nil {
		panic(err)
	}
	api := server.PathPrefix("/api").Subrouter()

	api.Use(middlewares.ContentTypeJson)
	config.InitApi(dbConn, api)

	api.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")

		res := Response{Message: "Anjing", Data: nil, Status: http.StatusOK}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(res)
	})
	// start server
	fmt.Println("application is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8888", server))
}
