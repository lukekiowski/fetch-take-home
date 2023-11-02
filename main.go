package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lukekiowski/fetchtakehome/pkg/handlers"
)

func main() {
    router := mux.NewRouter()

    // Here we'll define our api endpoints
    router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods(http.MethodPost)
    router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods(http.MethodGet)

    log.Println("API is running!")
    http.ListenAndServe(":8080", router)
}