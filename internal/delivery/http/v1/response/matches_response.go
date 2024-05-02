package response

type CreateMatch struct {
	MatchCatID string `json:"matchCatId"`
	UserCatID  string `json:"userCatId"`
	Message    string `json:"message"`
}
