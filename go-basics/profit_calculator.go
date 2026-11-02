package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, tax_rate float64

	fmt.Print("What is your total revenue? ")
	fmt.Scan(&revenue)

	fmt.Print("What are your total expenses? ")
	fmt.Scan(&expenses)

	fmt.Print("What is your current tax rate? ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit

	fmt.Printf("\nYour earnings before taxes (EBT) are: $%.2f\n", ebt)
	fmt.Printf("Your earnings after taxes are: $%.2f\n", profit)
	fmt.Printf("Your ratio of (EBT / profit) is: %.2f\n", ratio)
}
