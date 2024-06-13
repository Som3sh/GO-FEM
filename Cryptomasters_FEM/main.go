package main

import "sync"

func main() {
	currencies := []string{"btc", "eth", "bch"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()

}

//-----------------------------------------------------
// There are two methods to do so :
// 1. By putting it in the main.go file itself
// 2. Creating a new file and putting the function in that file [this makes the main.go file a bit cleaner and easier to modify the code so that we donot get confused on where is what in the same file]
//-----------------------------------------------------

// Code:
// func getCurrencyData(currency string) {

// 	rate, err := api.GetRate(currency)
// 	if err != nil {
// 		fmt.Println(err)

// 	} else {

// 		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
// 		fmt.Println("---------------------------")
// 	}
// }
