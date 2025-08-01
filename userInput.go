package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(usingCurrency *string, requiredCurrency *string, amount *float64) (err error) {
	for {
		fmt.Print("Введите вашу валюту: ")
		*usingCurrency, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		*usingCurrency = strings.TrimSpace(*usingCurrency)

		fmt.Print("Введите целевую валюту: ")
		*requiredCurrency, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		*requiredCurrency = strings.TrimSpace(*requiredCurrency)

		fmt.Print("Введите обменную сумму")
		amountStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		amountStr = strings.TrimSpace(amountStr)

		*amount, err = strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Введите числовое значение.")
			continue
		}
		return nil
	}
}
