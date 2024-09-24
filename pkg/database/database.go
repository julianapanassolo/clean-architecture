package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database representa a conexão com o banco de dados
type Database struct {
	DB *sqlx.DB
}

// NewDatabase cria uma nova instância de Database
func NewDatabase(dbURL string) (*Database, error) {
	db, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	return &Database{DB: db}, nil
}

// Close fecha a conexão com o banco de dados
func (db *Database) Close() error {
	return db.DB.Close()
}

// Exec executa uma consulta SQL sem retorno de dados
func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := db.DB.ExecContext(ctx, query, args...)
	return err
}

// Query executa uma consulta SQL com retorno de dados
func (db *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.QueryContext(ctx, query, args...)
}

// QueryRow executa uma consulta SQL com retorno de uma única linha
func (db *Database) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRowContext(ctx, query, args...)
}
