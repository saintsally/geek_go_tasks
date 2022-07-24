package repository

import (
	"context"
	"go_practice/go_practice/src/week4/apis/code"
	"go_practice/go_practice/src/week4/internal/domain"
	"go_practice/go_practice/src/week4/internal/repository/ent"

	"github.com/pkg/errors"
)

type repository struct {
	client *ent.Client
}

// NewRepository NewRepository
func NewRepository(client *ent.Client) domain.IUserRepo {
	return &repository{client: client}
}

// GetUserInfo GetUserInfo
func (r *repository) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	user, err := r.client.User.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, errors.Wrapf(code.UserNotFound, "user not found, id: %d, err: %+v", id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(code.Unknown, "db query err: %+v, id: %d,", err, id)
	}

	return &domain.User{
		Name: user.Name,
		City: user.City,
		ID:   user.ID,
	}, nil
}
