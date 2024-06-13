package main

func main() {
	getCurrencyData("btc")
	getCurrencyData("eth")
	getCurrencyData("bch")

}

//-----------------------------------------------------
// There are two methods to do so :
// 1. By putting it in the main.go file itself
// 2. Creating a new file and putting the function in that file [this makes the main.go file a bit cleaner and easier to modify the code so that we donot get confused on where is what in the same file]
//-----------------------------------------------------

// Code:
// func getCurrencyData() {
// 	var currency string
// 	var n int
// 	fmt.Print("How many currencies do you want to check? Enter : \t")

// 	fmt.Scanf("%d", &n)
// 	for i := 0; i < n; i++ {
// 		fmt.Print("\nEnter the currency you want to check : \t")
// 		fmt.Scanf("%s", &currency)

// 		rate, err := api.GetRate(currency)
// 		if err != nil {
// 			fmt.Println(err)

// 		} else {

// 			fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
// 			fmt.Println("---------------------------")
// 		}
// 	}
// 	fmt.Println("\n\n------------Thank You---------------")

// }
