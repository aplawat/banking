package main

import (
	"fmt"
)

type BankAccount struct {
	balance float64
}

type Bank struct {
	accounts map[string]*BankAccount
}

func (b *Bank) CreateAccount(accountID string) {
	b.accounts[accountID] = &BankAccount{balance: 0}
	fmt.Printf("Account %s created\n", accountID)
}

func (b *Bank) DeleteAccount(accountID string) {
	delete(b.accounts, accountID)
	fmt.Printf("Account %s deleted\n", accountID)
}

func (b *Bank) Deposit(accountID string, amount float64) {
	account, exists := b.accounts[accountID]
	if exists {
		account.balance += amount
		fmt.Printf("Deposited %.2f into account %s\n", amount, accountID)
	} else {
		fmt.Println("Account does not exist")
	}
}

func (b *Bank) Withdraw(accountID string, amount float64) {
	account, exists := b.accounts[accountID]
	if exists {
		if account.balance >= amount {
			account.balance -= amount
			fmt.Printf("Withdrawn %.2f from account %s\n", amount, accountID)
		} else {
			fmt.Println("Insufficient balance")
		}
	} else {
		fmt.Println("Account does not exist")
	}
}

func (b *Bank) DisplayBalance(accountID string) {
	account, exists := b.accounts[accountID]
	if exists {
		fmt.Printf("Account %s balance: %.2f\n", accountID, account.balance)
	} else {
		fmt.Println("Account does not exist")
	}
}

func (b *Bank) PrintAccounts() {
	fmt.Println("Accounts:")
	for accountID, account := range b.accounts {
		fmt.Printf("Account ID: %s, Balance: %.2f\n", accountID, account.balance)
	}
}

func main() {
	bank := &Bank{
		accounts: make(map[string]*BankAccount),
	}

	fmt.Println("Welcome")
	for {
		fmt.Println("Enter what you want to do:")
		fmt.Println("1 - Create account")
		fmt.Println("2 - Display balance")
		fmt.Println("3 - Deposit amount")
		fmt.Println("4 - Withdraw amount")
		fmt.Println("5 - Delete account")
		fmt.Println("0 - Exit")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		}

		var AccNo string
		fmt.Println("Enter your Account Number:")
		fmt.Scanln(&AccNo)

		switch choice {
		case 1:
			bank.CreateAccount(AccNo)
		case 2:
			bank.DisplayBalance(AccNo)
		case 3:
			fmt.Println("Enter the amount to deposit:")
			var amt float64
			fmt.Scanln(&amt)
			bank.Deposit(AccNo, amt)
		case 4:
			fmt.Println("Enter the amount to withdraw:")
			var amt float64
			fmt.Scanln(&amt)
			bank.Withdraw(AccNo, amt)
		case 5:
			bank.DeleteAccount(AccNo)
		default:
			fmt.Println("Invalid choice")
		}
	}
	bank.PrintAccounts()
}
