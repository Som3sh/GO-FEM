package main

import (
	"fmt"

	"learninggo.com/go/Cryptomasters/api"
)

func getCurrencyData(currency string) {

	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println(err)

	} else {

		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
		fmt.Println("---------------------------")
	}
}
