package request

type CreateCat struct {
	Name        string   `json:"name" validate:"required,min=5,max=20"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex"`
	AgeInMonth  int32    `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
}
