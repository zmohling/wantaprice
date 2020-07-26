package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
    "net/http"
)

type PriceResults struct {
	Name string `json:"Name"`
	Thumbnail string `json:"Thumbnail"`
	Price string `json:"Price"`
}

func getPrices(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
    http.HandleFunc(os.Getenv("BASE_URL") + "/scraper/getUPC", getUPC)
    http.HandleFunc(os.Getenv("BASE_URL") + "/scraper/getPrices", getUPC)
    log.Fatal(http.ListenAndServe(os.Getenv("SCRAPER_PORT"), nil))
}

func main() {
    handleRequests()
}