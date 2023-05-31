package main

import (
	"fmt"
	"math/rand"

	"time"
)

type BankAccount struct {
	balance int
}

func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
	fmt.Printf("Deposited %d, new balance is %d\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount int) {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d, new balance is %d\n", amount, a.balance)
	} else {
		fmt.Printf("Not enough funds to withdraw %d, balance is %d\n", amount, a.balance)
	}
}

func Problem5() {
	account := &BankAccount{balance: 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				if rand.Intn(2) == 0 {
					account.Deposit(rand.Intn(100))
				} else {
					account.Withdraw(rand.Intn(100))
				}
			}
		}()
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Printf("Current balance is %d\n", account.balance)
	}
}
