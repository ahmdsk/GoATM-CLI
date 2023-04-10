package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var choice int
	var amountToDeposit int
	var amountToWithdraw int

	var delayTime time.Duration = 3

	var balance int

	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println("Welcome to ATM System GoLang")
		fmt.Println("")
		fmt.Println("Please select an option")
		fmt.Println("")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Balance")
		fmt.Println("")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")

			fmt.Scan(&amountToDeposit)
			fmt.Println("You have deposited", amountToDeposit)
			balance += amountToDeposit

			fmt.Println("Your balance is", balance)

			time.Sleep(delayTime * time.Second)
		case 2:
			fmt.Print("Enter amount to withdraw: ")

			fmt.Scan(&amountToWithdraw)
			fmt.Println("You have withdrawn", amountToWithdraw)

			if amountToWithdraw > balance {
				fmt.Println("You don't have enough balance")
			} else {
				balance -= amountToWithdraw
			}

			fmt.Println("Your balance is", balance)

			time.Sleep(delayTime * time.Second)
		case 3:
			fmt.Println("Your balance is", balance)

			time.Sleep(delayTime * time.Second)
		default:
			fmt.Println("Invalid option")
			os.Exit(1)
		}
	}
}
