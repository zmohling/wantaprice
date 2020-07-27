package main

import (
    "fmt"
    "log"
		"github.com/anaskhan96/soup"
    "os"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type PriceResult struct {
	Name string `json:"Name"`
	Merchant string `json:"Merchant"`
	Price string `json:"Price"`
}

func getTopPrices(UPC string) []PriceResult {
	var url = "https://shopping.google.com/u/0/search?q=" + UPC
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	names := doc.FindAll("div", "class", "rgHvZc")
	var i int = 0

	var returnResults []PriceResult

	for _, name := range names {
		if(name.Find("a").Text() == "") {
			names = append(names[:i], names[i+1:]...)
		} else {
			i++
		}
	}

	i = 0
	prices := doc.FindAll("span", "class", "HRLxBb")
	for _, price := range prices {
		if(price.Text() == "") {
			prices = append(prices[:i], prices[i+1:]...)
		} else {
			i++
		}
	}

	i = 0
	merchants := doc.FindAll("div", "class", "dD8iuc")
	for _, merchant := range merchants {
		if(merchant.Text() == "") {
			merchants = append(merchants[:i], merchants[i+1:]...)
		} else {
			i++
		}
	}

	i = 0
	for range merchants {
		fmt.Println(names[i].Find("a").Text())
		fmt.Println(merchants[i].Text())
		fmt.Println(prices[i].Text())
		returnResults = append(returnResults, PriceResult{Name: names[i].Find("a").Text(), Merchant: merchants[i].Text(), Price: prices[i].Text()})
		i++
	}	
	fmt.Println(returnResults)
	return returnResults
}

func getUPC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	results := getTopPrices(vars["id"])
	for _, result := range results {
		json.NewEncoder(w).Encode(result)
	}
}

func getPrices(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "getPrices Endpoint")
	fmt.Println("")
}

func handleRequests() {
	  myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc(/*os.Getenv("BASE_URL") +*/ "/scraper/getUPC/{id}", getUPC)
    http.HandleFunc(/*os.Getenv("BASE_URL") +*/ "/scraper/getPrices", getPrices)
    log.Fatal(http.ListenAndServe(":8080"/*os.Getenv("SCRAPER_PORT")*/, myRouter))
}

func main() {
    handleRequests()
}