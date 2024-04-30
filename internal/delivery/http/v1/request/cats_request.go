package request

type CreateCat struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex" validate:"required"`
	AgeInMonth  int32    `json:"ageInMonth" validate:"required"`
	Description string   `json:"description" validate:"required"`
	ImageUrls   []string `json:"imageUrls" validate:"required,http_url"`
}

type GetCats struct {
	ID         string `query:"id" validate:"uuid"`
	Limit      int32  `query:"limit"`
	Offset     int32  `query:"offset"`
	Race       string `query:"race" validate:"oneof=Persian 'Maine Coon' Ragdoll Bengal Sphynx 'British Shorthair' Abyssinian 'Scottish Fold' Birman"`
	Sex        string `query:"sex" validate:"oneof=male female"`
	HasMatched bool   `query:"hasMatched" validate:"boolean"`
	AgeInMonth string `query:"ageInMonth"`
	Owned      bool   `query:"owned" validate:"boolean"`
	Search     string `query:"search"`
}
