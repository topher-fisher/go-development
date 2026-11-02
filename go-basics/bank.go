package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your balance is: $%.2f\n", accountBalance)
	} else if choice == 2 {
		fmt.Print("How much do you want to deposit? ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Balance updated! New balance amount: $%.2f\n", accountBalance)
	} else if choice == 3 {
		fmt.Print("How much money would you like to withdraw? ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Printf("Balance updated! New balance amount: $%.2f\n", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}

}
