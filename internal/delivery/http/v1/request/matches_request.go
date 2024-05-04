package request

type CreateMatch struct {
	MatchCatID string `json:"matchCatId" validate:"required"`
	UserCatID  string `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type ApproveMatch struct {
	MatchID string `json:"matchId" validate:"required"`
}

type DeleteMatch struct {
	MatchID string `json:"matchId" validate:"required"`
}
