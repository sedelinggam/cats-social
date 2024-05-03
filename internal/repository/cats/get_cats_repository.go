package catsRepository

import (
	"cats-social/internal/delivery/http/v1/request"
	"cats-social/internal/entity"
	"cats-social/pkg/util"
	"context"
	"fmt"
	"strings"
)

func (cr catRepository) GetCats(ctx context.Context, req request.GetCats) (*[]entity.Cat, error) {
	var (
		filterValues []interface{}
		whereClauses []string
		resp         []entity.Cat
		err          error
	)

	shouldFilters := ctx.Value("should_filters").(request.ShouldFilters)

	query := `SELECT
	id,
	user_id,
	name,
	race,
	sex,
	age_in_month,
	image_urls,
	description,
	is_already_matched,
	created_at
FROM
	cats`

	if shouldFilters.ID {
		filterValues = append(filterValues, req.ID)
		whereClauses = append(whereClauses, fmt.Sprintf("id = $%d", len(filterValues)))
	}

	if shouldFilters.Race {
		filterValues = append(filterValues, req.Race)
		whereClauses = append(whereClauses, fmt.Sprintf("race = $%d", len(filterValues)))
	}

	if shouldFilters.Sex {
		filterValues = append(filterValues, req.Sex)
		whereClauses = append(whereClauses, fmt.Sprintf("sex = $%d", len(filterValues)))
	}

	if shouldFilters.HasMatched {
		filterValues = append(filterValues, req.HasMatched)
		whereClauses = append(whereClauses, fmt.Sprintf("is_already_matched = $%d", len(filterValues)))
	}

	if shouldFilters.AgeInMonth {
		switch {
		case strings.HasPrefix(req.AgeInMonth, ">"):
			num := util.IsValidAge(req.AgeInMonth, ">")
			filterValues = append(filterValues, num)
			whereClauses = append(whereClauses, fmt.Sprintf("age_in_month >= $%d", len(filterValues)))
		case strings.HasPrefix(req.AgeInMonth, "<"):
			num := util.IsValidAge(req.AgeInMonth, "<")
			filterValues = append(filterValues, num)
			whereClauses = append(whereClauses, fmt.Sprintf("age_in_month <= $%d", len(filterValues)))
		default:
			num := util.IsValidAge(req.AgeInMonth, "")
			filterValues = append(filterValues, num)
			whereClauses = append(whereClauses, fmt.Sprintf("age_in_month = $%d", len(filterValues)))
		}
	}

	if shouldFilters.Owned {
		userID := ctx.Value("user_id").(string)
		filterValues = append(filterValues, userID)
		whereClauses = append(whereClauses, fmt.Sprintf("user_id = $%d", len(filterValues)))
	}

	if shouldFilters.Search {
		filterValues = append(filterValues, req.Search)
		whereClauses = append(whereClauses, fmt.Sprintf("LOWER(CONCAT_WS(' ', id, user_id, race, name, sex, description, age_in_month, image_urls, is_already_matched, created_at)) ILIKE '%%' || $%d || '%%'", len(filterValues)))
	}

	if len(whereClauses) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(whereClauses, " AND "))
	}

	query += " ORDER BY created_at ASC"

	filterValues = append(filterValues, req.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(filterValues))

	filterValues = append(filterValues, req.Offset)
	query += fmt.Sprintf(" OFFSET $%d", len(filterValues))

	err = cr.db.Select(&resp, query, filterValues...)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
