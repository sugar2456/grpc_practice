package repository

import (
	"context"

	"grpc_practice/internal/domain/model"
)

// UserRepository はユーザーデータのリポジトリインターフェース
type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}
