package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct { // this struct declare bank account and balce type as float64
	balance float64
}

func (b *BankAccount) Deposit(amount float64) { 
	b.balance += amount
	fmt.Println("Deposit successful")
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("Insufficient balance")
		return
	}
	b.balance -= amount
	fmt.Println("Withdrawal successful")
}

func (b *BankAccount) CheckBalance() {
	fmt.Printf("Current balance: %.2f\n", b.balance)
}

func main() {
	account := BankAccount{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Bank System ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("q. Quit")
		fmt.Print("Choose an option: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			// FIX: exit if no input (prevents timeout)
			fmt.Println("\nNo input detected. Exiting...")
			return
		}

		input = strings.TrimSpace(input)

		if input == "q" {
			fmt.Println("Exiting...")
			return
		}

		switch input {
		case "1":
			fmt.Print("Enter amount to deposit: ")
			val, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error. Exiting...")
				return
			}

			val = strings.TrimSpace(val)
			amount, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}
			account.Deposit(amount)

		case "2":
			fmt.Print("Enter amount to withdraw: ")
			val, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error. Exiting...")
				return
			}

			val = strings.TrimSpace(val)
			amount, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("Invalid amount")
				continue
			}
			account.Withdraw(amount)

		case "3":
			account.CheckBalance()

		default:
			fmt.Println("Invalid option")
		}
	}
}
