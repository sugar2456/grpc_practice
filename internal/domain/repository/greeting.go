package repository

import (
	"context"

	"grpc_practice/internal/domain/model"
)

// GreetingRepository は挨拶データのリポジトリインターフェース
type GreetingRepository interface {
	// Save は挨拶を保存する
	Save(ctx context.Context, greeting *model.Greeting) error
	// FindByID はIDで挨拶を検索する
	FindByID(ctx context.Context, id string) (*model.Greeting, error)
	// FindByName は名前で挨拶を検索する
	FindByName(ctx context.Context, name string) (*model.Greeting, error)
}
