package usersRepository

import (
	"cats-social/internal/entity"
	"context"
	"fmt"
)

func (ur userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var (
		resp *entity.User
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "email" = $1`, resp.TableName())
	err = ur.db.GetContext(ctx, &resp, query, email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
