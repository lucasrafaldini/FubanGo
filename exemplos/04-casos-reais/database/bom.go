package database

import (
	"context"
	"database/sql"
	"time"
)

// Uso correto: preparar statement e reutilizar conexões
func GoodQuery(db *sql.DB, name string) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "SELECT id, name FROM users WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, name)
}

// Transação segura com rollback em caso de erro
func SafeTransaction(db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO users(name) VALUES($1)", "x")
	if err != nil {
		return err
	}

	return nil
}
