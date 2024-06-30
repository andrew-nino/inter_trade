package entity

import "time"

type User struct {
	Id        int       `db:"id" swagg:"-"`
	Username  string    `db:"username"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at" swagg:"-"`
}

type StringToHash struct {
	String string `db:"string"`
}
