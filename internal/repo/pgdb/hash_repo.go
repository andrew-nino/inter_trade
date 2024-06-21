package pgdb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type HashToPostgres struct {
	db *sqlx.DB
}

func NewHashToPostgres(db *sqlx.DB) *HashToPostgres {
	return &HashToPostgres{db: db}
}

func (h *HashToPostgres) AddingHash(input string, hash string) (string, error) {

	var result string

	query := fmt.Sprintf("INSERT INTO %s (key, hash) VALUES ($1,$2) RETURNING created_at", hashTable)
	row := h.db.QueryRow(query, input, hash)
	if err := row.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
