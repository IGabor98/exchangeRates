package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Rates struct {
	Date  int    `json:"timestamp"`
	Base  string `json:"base"`
	Rates map[string]float64
}

const URL = "https://openexchangerates.org/api/latest.json?app_id=%s"

func main() {
	loadENV()

	apiKey := os.Getenv("API_KEY")
	res, err := http.Get(fmt.Sprintf(URL, apiKey))

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	rates, err := getRates(body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rates.Rates)

}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getRates(body []byte) (*Rates, error) {
	rates := new(Rates)
	err := json.Unmarshal(body, &rates)

	if err != nil {
		fmt.Println("whoops:", err)
	}

	return rates, err
}
