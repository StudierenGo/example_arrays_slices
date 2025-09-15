package main

import (
	"fmt"
)

func main() {
	transactions := []float64{}

	for {
		transaction := userTransaction()
		transactions = append(transactions, transaction)

		if transaction == 0 {
			break
		}
	}

	total := fmt.Sprintf("Ваш итоговый баланс равен: %.2f", totalTransactions(transactions))
	fmt.Println(total)
}

func userTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию (n для завершения): ")
	fmt.Scan(&transaction)

	return transaction
}

func totalTransactions(transactions []float64) float64 {
	total := 0.0
	for _, value := range transactions {
		total += value
	}

	return total
}
