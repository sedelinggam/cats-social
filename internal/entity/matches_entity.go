package entity

import "time"

type Match struct {
	ID         string    `db:"id"`
	MatchCatID string    `db:"match_cat_id"`
	UserCatID  string    `db:"user_cat_id"`
	Message    string    `db:"message"`
	CreatedAt  time.Time `db:"created_at"`
}

func (g Match) TableName() string {
	return `matches`
}
