package main

import (
	"math/rand"
	"time"
)

type TrasferRequest struct{
	ToAccunt int `json:"toaccout"`
	Amount int `json:"accout"`
}

type CreateAccountRequest struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Number    int       `json:"number"`
	Balance   int       `json:"balance"`
	CreateAt  time.Time `json:"createAt"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastname"`
	Number    int       `json:"number"`
	Balance   int       `json:"balance"`
	CreateAt  time.Time `json:"createAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Intn(1000000),
		Balance:   0,
		CreateAt:  time.Now().UTC(),
	}
}
