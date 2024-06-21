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

func (h *HashToPostgres) AddingHash(inputKey string, typeHash string, hash string) (string, error) {

	var result string

	query := fmt.Sprintf("INSERT INTO %s (type_hash, key, hash) VALUES ($1,$2, $3) RETURNING created_at", hashTable)
	row := h.db.QueryRow(query, typeHash, inputKey, hash)
	if err := row.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}

func (h *HashToPostgres) DeleteHash(inputKey, typeHash string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE key = $1 AND type_hash = $2", hashTable)

	_, err := h.db.Exec(query, inputKey, typeHash)

	return err
}
