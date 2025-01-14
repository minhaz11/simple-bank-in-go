package main

import "math/rand"

type Account struct {
	Id        int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AccNumber int64 `json:"acc_number"`
	Balance   float64 `json:"balance"`
}

func NewAccount(fname, lname string) *Account {
	return &Account{
		Id:        int64(rand.Intn(10000)),
		FirstName: fname,
		LastName:  lname,
		AccNumber: int64(rand.Intn(1000000000)),
		Balance: 0,
	}
}