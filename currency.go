package main

func currency() {
	// Курсы валют (1 базовая валюта = X целевой валюты)
	usdToRub := 91.50 // 1 USD = 91.50 RUB
	usdToYuan := 7.20 // 1 USD = 7.20 CNY
	usdToEur := 0.92  // 1 USD = 0.92 EUR

	yuanToRub := 12.71 // 1 CNY = 12.71 RUB (91.50/7.20)
	yuanToUsd := 0.14  // 1 CNY = 0.14 USD (1/7.20)
	yuanToEur := 0.13  // 1 CNY = 0.13 EUR

	eurToUsd := 1.08  // 1 EUR = 1.08 USD
	eurToRub := 98.82 // 1 EUR = 98.82 RUB (91.50*1.08)
	eurToYuan := 7.78 // 1 EUR = 7.78 CNY (7.20*1.08)

	rubToEur := 0.0101  // 1 RUB = 0.0101 EUR (1/98.82)
	rubToYuan := 0.0787 // 1 RUB = 0.0787 CNY (1/12.71)
	rubToUsd := 0.0109  // 1 RUB = 0.0109 USD (1/91.50)
}
