package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var allCurrency = []string{"Rub", "Eur", "Yuan", "Usd"}
	var usingCurrency, requiredCurrency string
	var amount float64
	isWork := true

	for isWork {
		printAllCurrency(allCurrency)
		errInput := userInput(&usingCurrency, &requiredCurrency, &amount)
		if errInput != nil {
			fmt.Println(errInput)
		}
		rate, errCurrency := rates[usingCurrency][requiredCurrency]
		if !errCurrency {
			fmt.Println("Курс не найден!")
		}
		convResult := amount * rate
		fmt.Println(convResult)
		for {
			fmt.Print("\nПовторить операцию?\n  (y -да)\n  (n - нет)")
			work, _ := bufio.NewReader(os.Stdin).ReadString('\n')

			switch work {
			case "y":
				continue
			case "n":
				return
			}
		}

	}
}
