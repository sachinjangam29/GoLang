package balance

import (
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile() (float64, error) {
	balanceText, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, fmt.Errorf("error in reading file: %w", err)
	}

	balance := string(balanceText)
	accountBalance, err := strconv.ParseFloat(balance, 64)

	if err != nil {
		return 0, fmt.Errorf("error in converting bytes to a string %w", err)
	}
	return accountBalance, nil
}

func WithdrawAmountBank() (float64, error) {
	var withdrawalAmount float64
	fmt.Println("Withdrawal Amount : ")
	fmt.Scan(&withdrawalAmount)

	if withdrawalAmount <= 0 {
		return 0, fmt.Errorf("incorrect number, The number cannot be or equal to 0")
	}

	balance, err := ReadBalanceFromFile()
	if err != nil {
		return 0, fmt.Errorf("error in the file %w ", err)
	}
	if withdrawalAmount > balance {
		return 0, fmt.Errorf("withdrawal amount exceeds your account balance")
	}

	balance -= withdrawalAmount
	writeBalanceToFile(balance)
	return balance, nil
}

func DepositAmountBank() (float64, error) {
	var depositAmount float64
	fmt.Println("Deposit Amount : ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		return 0, fmt.Errorf("incorrect number, the number cannot be or equal to 0")
	}
	balance, err := ReadBalanceFromFile()
	if err != nil {
		return 0, fmt.Errorf("error in the file %w", err)
	}
	balance += depositAmount
	fmt.Println("Your Account Balance is : ", balance)
	writeBalanceToFile(balance)
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
