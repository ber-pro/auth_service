package repository

import (
	"auth/internal/model"
	"context"
)

type Repository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.UserInfo) (int64, error)
	Update(ctx context.Context, id int64, user *model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}
