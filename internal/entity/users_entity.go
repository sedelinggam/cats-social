package entity

import "time"

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func (u User) TableName() string {
	return `users`
}

func (u User) NewEmail(email string) (string, error) {
	return "", nil
}

func (u User) NewPassword(password string) (string, error) {
	return "", nil
}
