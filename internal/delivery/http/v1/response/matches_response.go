package response

type CreateMatch struct {
	MatchCatID string `json:"matchCatId"`
	UserCatID  string `json:"userCatId"`
	Message    string `json:"message"`
}

type GetMatches struct {
	ID             string  `json:"id"`
	IssuedBy       GetUser `json:"issuedBy"`
	MatchCatDetail GetCat  `json:"matchCatDetail"`
	UserCatDetail  GetCat  `json:"userCatDetail"`
	Message        string  `json:"message"`
	CreatedAt      string  `json:"createdAt"`
}

type ApproveMatch struct {
	MatchID    string `json:"matchId"`
	UserCatID  string `json:"userCatId"`
	MatchCatID string `json:"matchCatId"`
}

type DeleteMatch struct {
	MatchID    string `json:"matchId"`
	UserCatID  string `json:"userCatId"`
	MatchCatID string `json:"matchCatId"`
}
