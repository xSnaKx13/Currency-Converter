package main

import (
	"fmt"
)

func main() {
	var allCurrency = []string{"Rub", "Eur", "Yuan", "Usd"}
	var usingCurrency, requiredCurrency string
	var amount float64

	for {
		printAllCurrency(allCurrency)
		errInput := userInput(&usingCurrency, &requiredCurrency, &amount)
		if errInput != nil {
			fmt.Println(errInput)
			continue
		}
		rate, errCurrency := rates[usingCurrency][requiredCurrency]
		if !errCurrency {
			fmt.Println("Курс не найден!\nПроверьте ввод валют")
			continue
		}
		convResult := amount * rate
		fmt.Printf("Ваш обменный курс %s к %s равен %.2f\n", usingCurrency, requiredCurrency, rate)
		fmt.Printf("Вы получили: %.2f\n", convResult)

		if !repiat() {
			return
		}
	}
}
