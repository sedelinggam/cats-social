package response

type CreateCat struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type GetCat struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Sex         string `json:"sex"`
	Description string `json:"description"`
	AgeInMonth  int32  `json:"ageInMonth"`
	ImageUrls   string `json:"imageUrls"`
	HasMatched  bool   `json:"hasMatched"`
	CreatedAt   string `json:"createdAt"`
}
