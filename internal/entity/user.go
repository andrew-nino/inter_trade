package entity

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
}

type StringToHash struct {
	String    string    `db:"string"`
	CreatedAt time.Time `db:"created_at"`
}

