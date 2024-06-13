package main

import (
	"fmt"

	"learninggo.com/go/Cryptomasters/api"
)

func getCurrencyData() {
	var currency string
	var n int
	fmt.Print("How many currencies do you want to check? Enter : \t")

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Print("\nEnter the currency you want to check : \t")
		fmt.Scanf("%s", &currency)

		rate, err := api.GetRate(currency)
		if err != nil {
			fmt.Println(err)

		} else {

			fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
			fmt.Println("---------------------------")
		}
	}
	fmt.Println("\n\n------------Thank You---------------")

}
