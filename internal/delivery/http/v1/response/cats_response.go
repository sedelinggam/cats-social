package response

type CreateCat struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type DeleteCat struct {
	ID        string `json:"id"`
	DeletedAt string `json:"deletedAt"`
}