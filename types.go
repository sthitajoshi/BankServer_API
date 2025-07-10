package main

import (
	"math/rand"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	Number    int    `json:"number"`
	Balance   int    `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(1000000),
		Balance:   0,
	}
}
