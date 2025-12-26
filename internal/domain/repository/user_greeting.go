package repository

import (
	"context"

	"grpc_practice/internal/domain/model"
)

// UserGreetingRepository は交差テーブルのリポジトリインターフェース
type UserGreetingRepository interface {
	// Save はユーザーと挨拶の関連を保存
	Save(ctx context.Context, ug *model.UserGreeting) error
	// FindByUserID はユーザーIDで関連を検索
	FindByUserID(ctx context.Context, userID string) ([]*model.UserGreeting, error)
	// FindByGreetingID は挨拶IDで関連を検索
	FindByGreetingID(ctx context.Context, greetingID string) ([]*model.UserGreeting, error)
	// FindGreetingsByUserID はユーザーIDで挨拶一覧を取得（JOINイメージ）
	FindGreetingsByUserID(ctx context.Context, userID string) ([]*model.Greeting, error)
	// FindUsersByGreetingID は挨拶IDでユーザー一覧を取得（JOINイメージ）
	FindUsersByGreetingID(ctx context.Context, greetingID string) ([]*model.User, error)
}
