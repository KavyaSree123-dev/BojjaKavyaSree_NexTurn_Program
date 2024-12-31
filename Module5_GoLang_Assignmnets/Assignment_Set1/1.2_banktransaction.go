// Bank Transaction Management System in go
package main

import (
	"errors"
	"fmt"
)

// Account represents a bank account with an ID, name, balance, and transaction log
// Each operation (deposit, withdraw) updates the transaction log
type Account struct {
	ID             int
	Name           string
	Balance        float64
	TransactionLog []string
}

// Bank maintains a collection of accounts and its details
type Bank struct {
	Name     string
	Address  string
	Accounts []Account
}

// Constants for menu options
const (
	OptionDeposit        = 1
	OptionWithdraw       = 2
	OptionCheckBalance   = 3
	OptionTransactionLog = 4
	OptionExitProgram    = 5
)

// Deposit increases the account balance and logs the transaction
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	a.TransactionLog = append(a.TransactionLog, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw decreases the account balance and logs the transaction
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	a.TransactionLog = append(a.TransactionLog, fmt.Sprintf("Withdrawn: %.2f", amount))
	return nil
}

// ShowTransactionLog displays the transaction history of an account
func (a *Account) ShowTransactionLog() {
	if len(a.TransactionLog) == 0 {
		fmt.Println("No transactions recorded.")
		return
	}
	fmt.Println("Transaction History:")
	for _, transaction := range a.TransactionLog {
		fmt.Println(transaction)
	}
}

func main() {
	bank := Bank{
		Name:    "Go Server Bank",
		Address: "456  Lane, Code City",
		Accounts: []Account{
			{ID: 1, Name: "Kavya", Balance: 2000.0},
			{ID: 2, Name: "Babu", Balance: 3000.0},
		},
	}

	var currentAccount *Account
	fmt.Printf("Welcome to %s\nAddress: %s\n", bank.Name, bank.Address)

	var accountID int
	fmt.Print("Enter your account ID: ")
	fmt.Scan(&accountID)
	for i := range bank.Accounts {
		if bank.Accounts[i].ID == accountID {
			currentAccount = &bank.Accounts[i]
			break
		}
	}
	if currentAccount == nil {
		fmt.Println("Invalid account ID.")
		return
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. View Transaction Log")
		fmt.Println("5. Exit Program")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case OptionDeposit:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := currentAccount.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successfully!")
			}
		case OptionWithdraw:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := currentAccount.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdraw successful!")
			}
		case OptionCheckBalance:
			fmt.Printf("Current Balance: %.2f\n", currentAccount.Balance)
		case OptionTransactionLog:
			currentAccount.ShowTransactionLog()
		case OptionExitProgram:
			fmt.Println("Thank you for using the Bank Transaction Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
