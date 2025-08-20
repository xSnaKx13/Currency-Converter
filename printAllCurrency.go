package main

import "fmt"

func printAllCurrency(allCurrency *[]string) {
	fmt.Println("\n--ДОСТУПНЫЕ ВАЛЮТЫ ДЛЯ ОБМЕНА--")
	for _, elem := range *allCurrency {
		fmt.Print(elem, " ")
	}
	fmt.Printf("\n\n")
}
