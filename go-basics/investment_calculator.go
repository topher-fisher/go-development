package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("What's Your Initial Investment Amount? ")
	fmt.Scan(&investmentAmount)

	fmt.Print("How Many Years Do You Plan On Investing? ")
	fmt.Scan(&years)

	fmt.Print("What's Your Expected Return Rate? ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
