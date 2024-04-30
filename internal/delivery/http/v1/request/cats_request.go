package request

type CreateCat struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex" validate:"required"`
	AgeInMonth  int32    `json:"ageInMonth" validate:"required"`
	Description string   `json:"description" validate:"required"`
	ImageUrls   []string `json:"imageUrls" validate:"required"`
}
