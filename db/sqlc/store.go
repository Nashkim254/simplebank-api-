package db

import (
	"context"
	"database/sql"
	"fmt"
)

//store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

//newStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//execTx executes a func within a database transaction
func (Store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err := fn(q)
	if err != nil {
		if rbErr := tx.RollBack(); rbErr != nil {
			return fmt.Errof("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
