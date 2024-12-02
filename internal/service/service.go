package service

import (
	desc "auth/pkg/user_v1"
	"context"
)

type Service interface {
	Get(ctx context.Context, id int64) (*desc.User, error)
	Create(ctx context.Context, user *desc.UserInfo) (int64, error)
	Update(ctx context.Context, id int64, user *desc.UserInfo) error
	Delete(ctx context.Context, id int64) error
}
