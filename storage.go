package main

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(id int64) error
	UpdateAccount(*Account) error
	GetAccountById(id int64) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error)  {
	connStr := "user=postgres dbname=postgres password=bank sslmode=disable"

	slog.Info("Opening connection to database")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	slog.Info("Connecting to database...")
	if err = db.Ping(); err != nil  {
		return nil, err
	}

	slog.Info("Connection to database successfull")
	return &PostgresStore{
		db: db,
	}, nil
}

func (ps *PostgresStore) CreateAccount(acc *Account) error {
	return nil
}

func (ps *PostgresStore) DeleteAccount(id int64) error {
	return nil
}

func (ps *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (ps *PostgresStore) GetAccountById(id int64) (*Account, error) {
	return nil,nil
}

