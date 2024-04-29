package request

type CreateCat struct {
	Name        string   `json:"name" validate:"required,min=5,max=20"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex"`
	AgeInMonth  int32    `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
}

type GetCats struct {
	ID               string `params:"id"`
	Limit            int32  `params:"limit"`
	Offset           int32  `params:"offset"`
	Race             string `params:"race"`
	Sex              string `params:"sex"`
	IsAlreadyMatched bool   `params:"isAlreadyMatched"`
	AgeInMonth       int32  `params:"ageInMonth"`
	Owned            bool   `params:"owned"`
	Search           string `params:"search"`
}
