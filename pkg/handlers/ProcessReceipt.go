package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/google/uuid"
    "github.com/lukekiowski/fetchtakehome/pkg/models"
    "github.com/lukekiowski/fetchtakehome/pkg/data"
    "fmt"
)


func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    // Read to request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var receipt models.Receipt
    json.Unmarshal(body, &receipt)

    // Append to the Receipt mocks
    receipt.Id = uuid.New().String()
    data.Receipts = append(data.Receipts, receipt)

    fmt.Println(data.Receipts)

    receiptCreatedResponse := models.ReceiptCreatedResponse{
        Id: receipt.Id,
    }

    // Send a 201 created response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(receiptCreatedResponse)
}