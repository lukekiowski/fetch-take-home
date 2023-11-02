package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "unicode/utf8"
    "github.com/gorilla/mux"
    "strconv"
    "strings"
    "math"
    "time"
    "github.com/lukekiowski/fetchtakehome/pkg/models"
    "github.com/lukekiowski/fetchtakehome/pkg/data"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {

    log.Println("GetPoints")
    vars := mux.Vars(r)
    id, ok := vars["id"]
    if !ok {
        log.Println("id is missing in parameters")
    }
    log.Println(`id := `, id)

    receipt := LookupReceipt(id)

    var points int = CalculatePoints(receipt)

    pointsResponse := models.PointsResponse {
        Points: points,
    }
    
    // Send a 200 response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(pointsResponse)

}

func LookupReceipt(id string)(models.Receipt){
    log.Println("Looking up", id)
    return *data.Receipts[id]
}

func CalculatePoints(receipt models.Receipt)(int){
    r, _ := json.Marshal(receipt)
    log.Println(string(r))
    var totalPoints = 0
    totalPoints += CountAlphaNumChars(receipt.Retailer)
    totalPoints += PointsForRoundDollar(receipt.Total)
    totalPoints += PointsForQuarterDollar(receipt.Total)
    totalPoints += PointsForEveryTwoItems(receipt.Items)
    totalPoints += PointsForItemsLength3(receipt.Items)
    totalPoints += PointsForOddDay(receipt.PurchaseDate)
    totalPoints += PointsFor2to4PM(receipt.PurchaseTime)
	return totalPoints
}

// One point for every alphanumeric character in the retailer name.
func CountAlphaNumChars(retailer string)(int){
    return utf8.RuneCountInString(retailer)
}

// 50 points if the total is a round dollar amount with no cents.
func PointsForRoundDollar(total string)(int){
    totalFloat, _ := strconv.ParseFloat(total, 64)
    var result = 0
    if totalFloat == float64(int64(totalFloat)){
        result = 50
    }
    log.Println("Adding", result, "points (50 points if the total is a round dollar amount with no cents)")
    return result
}

// 25 points if the total is a multiple of 0.25.
func PointsForQuarterDollar(total string)(int){
    totalFloat, _ := strconv.ParseFloat(total, 64)
    var result = 0
    if math.Mod(totalFloat, 0.25) == 0.0{
        result = 25
    }
    log.Println("Adding", result, "points (25 points if the total is a multiple of 0.25)")
    return result
}

// 5 points for every two items on the receipt.
func PointsForEveryTwoItems(items []models.Item)(int){
    var numItems = len(items)
    var result = 5 * int(math.Floor(float64(numItems)/2.0))
    log.Println("Adding", result, "points (5 points for every two items on the receipt)")
    return result
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. 
// The result is the number of points earned.
func PointsForItemsLength3(items []models.Item)(int){
    var pointsSoFar = 0
    for _, v := range items {
        pointsSoFar += PointsForDescriptionLength(v.ShortDescription, v.Price)
    } 
    log.Println("Adding", pointsSoFar, "points (If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer)")
    return pointsSoFar
}
func PointsForDescriptionLength(description string, price string)(int){
    priceFloat, _ := strconv.ParseFloat(price, 64)
    var trimmed = strings.Trim(description, " ")
    var length = len(trimmed)
    if length % 3 == 0{
        return int(math.Ceil(priceFloat * 0.2))
    }
    return 0
}

// 6 points if the day in the purchase date is odd.
func PointsForOddDay(dateStr string)(int){
    date, _ := time.Parse("2006-01-02", dateStr)
    var result = 0
    if date.Day() % 2 != 0{ // odd
        result = 6
    }
    log.Println("Adding", result, "points (6 points if the day in the purchase date is odd)")
    return result
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func PointsFor2to4PM(timeStr string)(int){
    var timeFormat = "15:04:05"
    timeParsed, _ := time.Parse(timeFormat, timeStr+":00")
    var twoPM, _ = time.Parse(timeFormat, "14:00"+":00") // 2pm
    var fourPM, _ = time.Parse(timeFormat, "16:00"+":00") // 4pm
    var result = 0
    if timeParsed.After(twoPM) && timeParsed.Before(fourPM){ 
        result = 10
    }
    log.Println("Adding", result, "points (6 points if the day in the purchase date is odd)")
    return result
}