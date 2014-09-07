package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Products hold store with product
type Products struct {
	Stores []Store `json:"stores"`
}

//Store with product
type Store struct {
	StoreID int    `json:"storeId"`
	Name    string `json:"name"`
}

var zip int
var distance int
var sku int
var apiKey string

func main() {
	flag.IntVar(&zip, "zip", 14214, "zip to find product")
	flag.IntVar(&distance, "distance", 50, "distance from zip to find product")
	flag.IntVar(&sku, "sku", 8307143, "sku of find product")
	flag.StringVar(&apiKey, "apiKey", "", "api key of best buy api")
	flag.Parse()
	if apiKey == "" {
		fmt.Println("best-buy-finder usage: ")
		flag.PrintDefaults()
		log.Fatal("Please input api key")
	}
	url := fmt.Sprintf("http://api.remix.bestbuy.com/v1/stores(area(%d,%d))+products"+
		"(sku=%d)?format=json&show=storeId,name,products.sku,products.name&"+
		"apiKey=%s", zip, distance, sku, apiKey)
	resp, _ := http.Get(url)
	var products Products
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &products)
	if len(products.Stores) == 0 {
		fmt.Println("No stores with sku")
	} else {
		fmt.Printf("Stores with Sku %d\n", sku)
		for _, value := range products.Stores {
			fmt.Printf("Store Id: %d, Store Name: %s\n", value.StoreID, value.Name)
		}
	}
}
