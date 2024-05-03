package matchesRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (mr matchRepository) GetMatches(ctx context.Context, userID string) ([]entity.Match, error) {
	var (
		resp []struct {
			entity.Match
			UserCat  entity.Cat  `db:"user_cat"`
			MatchCat entity.Cat  `db:"match_cat"`
			Issuer   entity.User `db:"issuer"`
		}
		err error
	)

	query := fmt.Sprintf(`
	SELECT 
			m.*,
			uc.id AS "user_cat.id",
			uc.name AS "user_cat.name",
			uc.race AS "user_cat.race",
			uc.sex AS "user_cat.sex",
			uc.description AS "user_cat.description",
			uc.age_in_month AS "user_cat.age_in_month",
			uc.image_urls AS "user_cat.image_urls",
			uc.is_already_matched AS "user_cat.is_already_matched",
			uc.created_at AS "user_cat.created_at",
			mc.id AS "match_cat.id", 
			mc.name AS "match_cat.name",
			mc.race AS "match_cat.race",
			mc.sex AS "match_cat.sex",
			mc.description AS "match_cat.description",
			mc.age_in_month AS "match_cat.age_in_month",
			mc.image_urls AS "match_cat.image_urls",
			mc.is_already_matched AS "match_cat.is_already_matched",
			mc.created_at AS "match_cat.created_at",
			u.id AS "issuer.id",
			u.name AS "issuer.name",
			u.email AS "issuer.email",
			u.created_at AS "issuer.created_at"
		FROM %s m
		LEFT JOIN %s uc ON m.user_cat_id = uc.id
		LEFT JOIN %s mc ON m.match_cat_id = mc.id
		LEFT JOIN %s u ON m.issuer_id = u.id
		WHERE m.issuer_id = $1 OR m.receiver_id = $1
	`, entity.Match{}.TableName(), entity.Cat{}.TableName(), entity.Cat{}.TableName(), entity.User{}.TableName())

	err = mr.db.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}

	var matches []entity.Match
	for _, r := range resp {
		match := r.Match
		match.UserCat = &r.UserCat
		match.MatchCat = &r.MatchCat
		match.Issuer = &r.Issuer
		matches = append(matches, match)
	}

	return matches, nil
}
