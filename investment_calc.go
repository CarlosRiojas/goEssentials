package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func calculateProfitFinancials(earnings float64, expenses float64) (float64, float64) {
	//hardvalues
	tax := .17

	//earnigs calculations
	earningsAfterTax := (earnings - tax*earnings) - expenses
	ratio := earnings / earningsAfterTax

	//write the output file here
	writeProfitFinancials := fmt.Sprint("Financials Report\n",
		"Earnings After Tax:$ ", earningsAfterTax, "\n", "Ratio:", ratio, "%")
	os.WriteFile("ProfitReport.txt", []byte(writeProfitFinancials), 0644)
	return ratio, earningsAfterTax
}

func calculateInvestmentFinancials(investmentAmount float64, years float64) (float64, float64) {
	expectedReturnRate := 5.5
	inflationRate := .05

	//value calculations
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//Write the output file here
	writeInvestmentFinancials := fmt.Sprint("Investments Report\n",
		"Future Value is: $", futureValue, "\n", "Future Real Value: $", futureRealValue, "\n",
		"At these rates \n", "Expected Rate: ", expectedReturnRate, "% \n", "Inflation Rate: ", (inflationRate * 100), "%\n")
	os.WriteFile("InvestmentReport.txt", []byte(writeInvestmentFinancials), 0644)
	return futureValue, futureRealValue
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return balance, nil
}

func readInvestmentStatement() {
	data, _ := os.Open("ProfitReport.txt")
	defer data.Close()
	buffer := make([]byte, 1024)
	n, err := data.Read(buffer)

	for err == nil {
		//Do something with the read data(slice of bytes)
		fmt.Println(string(buffer[:n]))
		n, err = data.Read(buffer)
	}
	if err != io.EOF {
		fmt.Println("End of Function", err)
	}
}

const accountBalanceFile = "ProfitReport.txt"
const investmentReportFile = "InvestmentReport.txt"

func main() {
	var accountBalance, _ = getFloatFromFile(accountBalanceFile)
	var investmentReport, _ = getFloatFromFile(investmentReportFile)

	var earnings float64
	var expenses float64
	var investmentAmount float64
	var years float64
	/*
		const inflationRate = 2.5
		investmentAmount := 1000
		expectedReturnRate := 5.5
		years := 10.0
	*/
	fmt.Println("Go Profit calculator!")
	fmt.Println("What would you like to do today?:")
	fmt.Println("1.-Expenses Mode")
	fmt.Println("2.-Investments Mode")
	fmt.Println("3.-Read Profit Statement")
	fmt.Println("4.-Read Investment Statement")
	fmt.Println("5.-Exit")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("Profit and Expenses System!\n")
		fmt.Print("Please add your earnings: $")
		fmt.Scan(&earnings)

		fmt.Print("Please add your expenses: $")
		fmt.Scan(&expenses)

		fmt.Print("Your profit calculations show as follow")
		ratio, profit := calculateProfitFinancials(earnings, expenses)
		fmt.Print("Your profit is: $", profit)
		fmt.Printf(".\n")
		fmt.Print("Your ratio is:", ratio, " %")
		fmt.Printf(".\n")
		fmt.Printf("Your report has been written.\n")
		fmt.Printf("\n")
		main()

	case 2:
		fmt.Printf("Investments System!\n")
		fmt.Print("Please add your investment Amount: $")
		fmt.Scan(&investmentAmount)

		fmt.Print("Please add the amount of years: ")
		fmt.Scan(&years)
		futureValue, futureRealValue := calculateInvestmentFinancials(investmentAmount, years)

		fmt.Print("Your future value is: $", futureValue, " at ", years, " years")
		fmt.Printf(".\n")
		fmt.Print("Your future real value is: $", futureRealValue)
		fmt.Printf(".\n")
		fmt.Printf("Your report has been written.\n")
		fmt.Printf("\n")
		main()
	case 3:
		fmt.Print(accountBalance)
		main()
	case 4:
		fmt.Print(investmentReport)
		main()
	case 5:
		fmt.Print("Bye!")
		break
	}

}
