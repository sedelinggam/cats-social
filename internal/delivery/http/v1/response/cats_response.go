package response

type CreateCat struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type UpdateCat struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updatedAt"`
}
