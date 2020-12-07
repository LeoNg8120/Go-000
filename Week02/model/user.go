package model


type User struct {
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Age int `db:"age"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}