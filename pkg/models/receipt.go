package models
// import (
//     "github.com/lukekiowski/fetchtakehome/pkg/models"
// )

type Receipt struct {
    Id string `json:"id"`
    Retailer string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Total string `json:"total"`
    Items []Item `json:"items"`
}