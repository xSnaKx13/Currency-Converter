package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func userInput(usingCurrency string, amount float64, requiredCurrency string) (string, float64, string, error) {

	printAllCurrency()
	fmt.Print("Выберете обменную валюту: ")
	usingCurrency, err := bufio.NewReader(os.Stdin).ReadString('\n')
	usingCurrency = strings.TrimSpace(usingCurrency)
	strings.ToLower(usingCurrency)
	if err != nil {
		return "", 0, "", errors.New("Такой валюты нет!")
	}

	fmt.Print("Колличество, которое хотите обменять: ")
	fmt.Scan(&amount)

	for _, elem := range allCurrency{
		if usingCurrency == strings.ToLower(elem){
			continue
		}
		fmt.Println(elem)
	}
	fmt.Println("Выберете целевую валюту: ")
	requiredCurrency, err = bufio.NewReader(os.Stdin).ReadString('\n')
	requiredCurrency = strings.TrimSpace(requiredCurrency)
	strings.ToLower(requiredCurrency)
	if err != nil {
		return "", 0, "", errors.New("Такой валюты нет!")
	}

	return usingCurrency, amount, requiredCurrency, nil
}
