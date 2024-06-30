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

// Adding a new record to the table. If successful, we return a string with the time the record was created.
// Or we get the [ErrNoRows] error if the entry was not added.
func (h *HashToPostgres) AddingHash(inputKey string, typeHash string, hash string) (string, error) {

	var result string

	query := fmt.Sprintf("INSERT INTO %s (type_hash, key, hash) VALUES ($1,$2, $3) RETURNING created_at", hashTable)
	row := h.db.QueryRow(query, typeHash, inputKey, hash)
	if err := row.Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}

// We delete a record from the table or get an error.
func (h *HashToPostgres) DeleteHash(inputKey, typeHash string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE key = $1 AND type_hash = $2", hashTable)

	_, err := h.db.Exec(query, inputKey, typeHash)

	return err
}
