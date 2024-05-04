package request

type CreateCat struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex" validate:"required"`
	AgeInMonth  int32    `json:"ageInMonth" validate:"required"`
	Description string   `json:"description" validate:"required"`
	ImageUrls   []string `json:"imageUrls" validate:"required"`
}

type UpdateCat struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Race        string   `json:"race" validate:"required"`
	Sex         string   `json:"sex" validate:"required"`
	AgeInMonth  int32    `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string   `json:"description" validate:"required"`
	ImageUrls   []string `json:"imageUrls" validate:"required"`
}
type GetCats struct {
	ID         string `query:"id" validate:"omitempty,uuid"`
	Limit      int32  `query:"limit"`
	Offset     int32  `query:"offset"`
	Race       string `query:"race" validate:"omitempty,oneof=Persian 'Maine Coon' Ragdoll Bengal Sphynx 'British Shorthair' Abyssinian 'Scottish Fold' Birman"`
	Sex        string `query:"sex" validate:"omitempty,oneof=male female"`
	HasMatched bool   `query:"hasMatched" validate:"omitempty,boolean"`
	AgeInMonth string `query:"ageInMonth"`
	Owned      bool   `query:"owned" validate:"omitempty,boolean"`
	Search     string `query:"search"`
}

type ShouldFilters struct {
	ID         bool
	Limit      bool
	Offset     bool
	Race       bool
	Sex        bool
	HasMatched bool
	AgeInMonth bool
	Owned      bool
	Search     bool
}
