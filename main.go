package main

import "fmt"

var allCurrency = []string{"Rub", "Eur", "Yuan", "Usd"}
var usingCurrency, requiredCurrency string
var amount float64

func main() {
	for {
		fmt.Println("Введите имеющуюся валюту")
		userInput(usingCurrency, amount, requiredCurrency)
	}
}
