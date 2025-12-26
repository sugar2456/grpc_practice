package usecase

import (
	"context"

	"grpc_practice/internal/domain/model"
	"grpc_practice/internal/domain/repository"
)

// GreeterUsecase は挨拶のユースケース
type GreeterUsecase interface {
	SayHello(ctx context.Context, name string) (*model.Greeting, error)
	SayHelloStream(ctx context.Context, name string, count int) ([]*model.Greeting, error)
}

type greeterUsecase struct {
	repo repository.GreetingRepository
}

// NewGreeterUsecase は新しいGreeterUsecaseを作成
func NewGreeterUsecase(repo repository.GreetingRepository) GreeterUsecase {
	return &greeterUsecase{repo: repo}
}

func (u *greeterUsecase) SayHello(ctx context.Context, name string) (*model.Greeting, error) {
	greeting := model.NewGreeting(name)

	// リポジトリに保存（オプション）
	if err := u.repo.Save(ctx, greeting); err != nil {
		return nil, err
	}

	return greeting, nil
}

func (u *greeterUsecase) SayHelloStream(ctx context.Context, name string, count int) ([]*model.Greeting, error) {
	greetings := make([]*model.Greeting, count)
	for i := 0; i < count; i++ {
		greetings[i] = &model.Greeting{
			Name:    name,
			Message: "Hello " + name + "!",
		}
	}
	return greetings, nil
}
