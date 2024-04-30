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
	ID         string `query:"id" validate:"uuid"`
	Limit      int32  `query:"limit" validate:"isdefault=5"`
	Offset     int32  `query:"offset" validate:"isdefault=0"`
	Race       string `query:"race" validate:"oneof=Persian"`
	Sex        string `query:"sex" validate:"oneof=male female"`
	HasMatched bool   `query:"hasMatched" validate:"boolean"`
	AgeInMonth string `query:"ageInMonth"`
	Owned      bool   `query:"owned" validate:"boolean"`
	Search     string `query:"search"`
}
