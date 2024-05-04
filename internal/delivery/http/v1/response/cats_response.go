package response

import (
	"time"
)

type CreateCat struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type UpdateCat struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
}
type GetCat struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	Description string   `json:"description"`
	AgeInMonth  int32    `json:"ageInMonth"`
	ImageUrls   []string `json:"imageUrls"`
	HasMatched  bool     `json:"hasMatched"`
	CreatedAt   string   `json:"createdAt"`
}

type GetCats struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int32     `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	Description string    `json:"description"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}
type DeleteCat struct {
	ID        string `json:"id"`
	DeletedAt string `json:"deletedAt"`
}
