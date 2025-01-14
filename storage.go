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

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=bank sslmode=disable"

	slog.Info("Opening connection to database")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	slog.Info("Connecting to database...")
	if err = db.Ping(); err != nil {
		return nil, err
	}

	slog.Info("Connection to database successfull")
	return &PostgresStore{
		db: db,
	}, nil
}

func (ps *PostgresStore) Seeder() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		account_number TEXT,
		balance DOUBLE PRECISION DEFAULT 0,
		created_at timestamp
	)`

	_, err := ps.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO accounts (first_name, last_name, account_number, balance, created_at) values ($1, $2, $3, $4, $5)`

	stmt, err := ps.db.Prepare(query)

	if err != nil {
		slog.Info("Error preparing insert query")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(acc.FirstName, acc.LastName, acc.AccNumber, acc.Balance, acc.CreatedAt)

	if err != nil {
		slog.Info("Error executing insert query")
		return err
	}

	return nil

}

func (ps *PostgresStore) DeleteAccount(id int64) error {
	return nil
}

func (ps *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (ps *PostgresStore) GetAccountById(id int64) (*Account, error) {
	return nil, nil
}
