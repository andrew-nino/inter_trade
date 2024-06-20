package entity

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type StringToHash struct {
	ID_User   int       `db:"id_user"`
	String    string    `db:"string"`
	CreatedAt time.Time `db:"created_at"`
}
