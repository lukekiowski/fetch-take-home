package data

import "github.com/lukekiowski/fetchtakehome/pkg/models"

var Receipts = []models.Receipt{
    {
        Id:     "111",
        Retailer:  "Golang",
        PurchaseDate: "Gopher",
        PurchaseTime:   "Gopher",
        Total:   "Gopher",
        Items:   []models.Item{{
            ShortDescription:	"Pepsi - 12-oz",
            Price:				"1.25",
        }},
    },
}
