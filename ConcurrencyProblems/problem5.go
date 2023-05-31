package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Problem 5

// BankAccount is a struct that represents a bank account which has a balance and a mutex
// the mutex is used so that we can lock the balance when required as we are using multiple go routines to access the balance
// if we do not lock the balance then the balance will be inconsistent as there can be deposits and withdrawals at the same time
// so it is necessary to lock the balance when we are accessing it to update it
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// when the execution of this function starts it locks the balance so that no other go routine can access it. So the balance will be consistent.
func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	fmt.Printf("Deposited %d, new balance is %d\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d, new balance is %d\n", amount, a.balance)
	} else {
		fmt.Printf("Not enough funds to withdraw %d, balance is %d\n", amount, a.balance)
	}
}

func Problem5() {
	account := &BankAccount{balance: 0}

	//rand seed is deprecated in go 1.20 and above . there is a alternative for it
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
		}() //there is a syntax error in the code provided in the problem statement here () is missing
	}

	// as the code provided in the problem statement the last loop keeps on running until some interrupt is provided
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Printf("Current balance is %d\n", account.balance)
	}

}
