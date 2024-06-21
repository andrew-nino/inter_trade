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

func (h *HashToPostgres) AddingHash(input string, typeHash string, hash string) (string, error) {

	var result string

	query := fmt.Sprintf("INSERT INTO %s (type_hash, key, hash) VALUES ($1,$2, $3) RETURNING created_at", hashTable)
	row := h.db.QueryRow(query, typeHash, input, hash)
	if err := row.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}
