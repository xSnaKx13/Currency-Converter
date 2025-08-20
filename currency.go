package main

func currency(rates *map[string]map[string]float64){

*rates = map[string]map[string]float64{
	"rub": {
		"usd":  0.0109,
		"eur":  0.0101,
		"yuan": 0.0787,
	},
	"usd": {
		"rub":  91.50,
		"eur":  0.92,
		"yuan": 7.20,
	},
	"yuan": {
		"rub": 12.71,
		"usd": 0.14,
		"eur": 0.13,
	},
	"eur": {
		"rub":  98.82,
		"usd":  1.08,
		"yuan": 7.78,
	},
}
}