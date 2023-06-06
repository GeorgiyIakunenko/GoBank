package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string, balance int) *Account {
	return &Account{
		ID:        rand.Intn(100),
		FirstName: firstName,
		LastName:  lastName,
		Balance:   int64(balance),
	}
}
