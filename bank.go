package main

import (
	bal "bank/balance"
	opt "bank/presentor"
	"fmt"
)

func main() {
	var choice int
	//var accountBalance float64 = 1000

	fmt.Println("Welcome to the Bank...")

	for {
		opt.PresentOptions()
		fmt.Scan(&choice)
		if choice == 1 {
			initialBalance, err := bal.ReadBalanceFromFile()
			if err != nil {
				fmt.Println("Error in the file: ", err)
				return
			}
			fmt.Println("the Account Balance is: ", initialBalance)
		} else if choice == 2 {

			depositAmount, err := bal.DepositAmountBank()
			if err != nil {
				fmt.Println("Error in processing transaction: ", err)
				continue
			}
			fmt.Println("The Total Balance is: ", depositAmount)

		} else if choice == 3 {
			withdrawalAmount, err := bal.WithdrawAmountBank()

			if err != nil {
				fmt.Println("Error in processing the transaction: ", err)
			}

			fmt.Println("The remaining balance is: ", withdrawalAmount)
		} else {
			fmt.Println("GoodBye !")
			break
		}
	}
	fmt.Println("Thank you for visiting the bank...")
}
