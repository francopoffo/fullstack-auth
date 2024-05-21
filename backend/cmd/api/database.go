package api

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

type Storage interface {
	Close() error
	Init() error
}

func NewPostgresDB() (*PostgresDB, error) {
	dataSourceName := "postgres://postgres:postgres@database:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db: db}, nil
}

func (s *PostgresDB) Close() error {
	return s.db.Close()
}

func (s *PostgresDB) Init() error {
	return s.CreateUserTable()
}

func (s *PostgresDB) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`
	_, err := s.db.Exec(query)
	return err
}
