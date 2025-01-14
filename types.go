package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct{
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type Account struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	AccNumber int64     `json:"acc_number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(fname, lname string) *Account {
	return &Account{
		Id:        int64(rand.Intn(10000)),
		FirstName: fname,
		LastName:  lname,
		AccNumber: int64(rand.Intn(1000000000)),
		Balance:   0,
		CreatedAt: time.Now(),
	}
}
