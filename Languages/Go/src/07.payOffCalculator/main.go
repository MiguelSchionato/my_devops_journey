package main

import (
	"fmt"
)

func main() {
	var totalDebt, monthlyPayment, monthlyInterest, totalInterest, totalPaid, initialDebt float32
	var month int

	fmt.Print("Total amount of debt: ")
	fmt.Scanf("%f\n", &totalDebt)

	fmt.Print("Monthly payment: ")
	fmt.Scanf("%f\n", &monthlyPayment)

	fmt.Print("Monthly interest: ")
	fmt.Scanf("%f\n", &monthlyInterest)

	initialDebt = totalDebt
	firstInterest := totalDebt * (monthlyInterest / 100)

	if monthlyPayment <= firstInterest {
		fmt.Println("Attention, monthly payment is less than the interest, you'll never pay off your debt with that rate.")
		return
	}

	for totalDebt > 0 {

		month++
		interestCharged := totalDebt * monthlyInterest / 100
		totalDebt = totalDebt - monthlyPayment + interestCharged
		totalInterest += interestCharged

		currentMonthOwed := totalDebt + interestCharged
		if currentMonthOwed <= monthlyPayment { // fix negative last payment forcing the variable to 0
			totalDebt = 0
		}

		fmt.Printf("%d° month: %.2f\n", month, totalDebt)
	}

	totalPaid = totalInterest + initialDebt
	fmt.Printf("You'll pay off your debt in %d months\n", month)
	fmt.Printf("Total interest paid: %.2f\n", totalInterest)
	fmt.Printf("Total paid: %.2f", totalPaid)
}
