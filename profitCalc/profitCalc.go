package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	const taxRate = 14.5

	fmt.Print("Please enter this month revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter this month expenses: ")
	fmt.Scan(&expenses)

	earningsBeforeTax := revenue - expenses

	profit := (earningsBeforeTax * (1 - taxRate/100))

	returnRatio := (earningsBeforeTax / profit)

	fmt.Println("Your Earnings before int is: $", earningsBeforeTax, " usd")
	fmt.Printf("Your profit is: $ %.1f usd\n", profit)
	fmt.Printf("Your returnRatio is:%.1f", returnRatio)
}
