package entity

import (
	"time"
)

type Cat struct {
	ID             string    `db:"id"`
	UserID         string    `db:"user_id"`
	User           User      `db:"user"`
	Name           string    `db:"name"`
	Race           string    `db:"race"`
	Sex            string    `db:"sex"`
	AgeInMonth     int32     `db:"age_in_month"`
	Image          string    `db:"image_urls"`
	Description    string    `db:"description"`
	IsAlreadyMatch bool      `db:"is_already_matched"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}

func (g Cat) TableName() string {
	return `cats`
}
