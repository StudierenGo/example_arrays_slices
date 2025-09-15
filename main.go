package main

import (
	"fmt"
	"strings"
)

func main() {
	transactions := []float64{}

	for {
		transaction := userTransaction()
		transactions = append(transactions, transaction)

		if !userAction() || transaction == 0 {
			break
		}
	}

	fmt.Println("Ваши транзакции: ", transactions)
}

func userTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию: ")
	fmt.Scan(&transaction)

	return transaction
}

func userAction() bool {
	var action string
	fmt.Println("Хотите продолжить? (y/n) ")
	fmt.Scan(&action)

	return strings.ToLower(action) == "y" || strings.ToLower(action) == "yes"
}
