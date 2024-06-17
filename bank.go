package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		var choice int
		var accountBalance float64 = 1000
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("Your Account Balance is : ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Println("Deposit Amount : ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Incorrect Number, The number cannot be or equal to 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your Account Balance is : ", accountBalance)
		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Println("Withdrawal Amount : ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Incorrect Number, The number cannot be or equal to 0")
				continue
			}

			if withdrawalAmount >= accountBalance {
				fmt.Println("Withdrawal Amount exceeds your Account Balance...")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your Account Balance is : ", accountBalance)
		} else {
			fmt.Println("GoodBye !")
			break
		}
	}
	fmt.Println("Thank you for visiting the bank...")
}
