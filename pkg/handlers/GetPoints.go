package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lukekiowski/fetchtakehome/pkg/models"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {

    log.Println("GetPoints")
    vars := mux.Vars(r)
    id, ok := vars["id"]
    if !ok {
        log.Println("id is missing in parameters")
    }
    log.Println(`id := `, id)

    var points int = 0 

    pointsResponse := models.PointsResponse {
        Points: points,
    }
    
    // Send a 200 response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(pointsResponse)

}