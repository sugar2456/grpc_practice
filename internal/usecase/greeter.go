package usecase

import (
	"context"
	"fmt"

	"grpc_practice/internal/domain/model"
	"grpc_practice/internal/domain/repository"
)

// GreeterUsecase は挨拶のユースケース
type GreeterUsecase interface {
	SayHello(ctx context.Context, userID, name string) (*model.Greeting, error)
	SayHelloStream(ctx context.Context, name string, count int) ([]*model.Greeting, error)
	GetUserGreetings(ctx context.Context, userID string) ([]*model.Greeting, error)
}

type greeterUsecase struct {
	greetingRepo     repository.GreetingRepository
	userRepo         repository.UserRepository
	userGreetingRepo repository.UserGreetingRepository
}

// NewGreeterUsecase は新しいGreeterUsecaseを作成
func NewGreeterUsecase(
	greetingRepo repository.GreetingRepository,
	userRepo repository.UserRepository,
	userGreetingRepo repository.UserGreetingRepository,
) GreeterUsecase {
	return &greeterUsecase{
		greetingRepo:     greetingRepo,
		userRepo:         userRepo,
		userGreetingRepo: userGreetingRepo,
	}
}

func (u *greeterUsecase) SayHello(ctx context.Context, userID, name string) (*model.Greeting, error) {
	// ユーザーを取得または作成
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user = model.NewUser(userID, name)
		if err := u.userRepo.Save(ctx, user); err != nil {
			return nil, err
		}
	}

	// 挨拶を作成
	greetingID := fmt.Sprintf("greeting-%s-%d", userID, len(name))
	greeting := model.NewGreeting(greetingID, name)
	if err := u.greetingRepo.Save(ctx, greeting); err != nil {
		return nil, err
	}

	// 交差テーブルに関連を保存
	userGreetingID := fmt.Sprintf("ug-%s-%s", userID, greetingID)
	userGreeting := model.NewUserGreeting(userGreetingID, userID, greetingID)
	if err := u.userGreetingRepo.Save(ctx, userGreeting); err != nil {
		return nil, err
	}

	return greeting, nil
}

func (u *greeterUsecase) SayHelloStream(ctx context.Context, name string, count int) ([]*model.Greeting, error) {
	greetings := make([]*model.Greeting, count)
	for i := 0; i < count; i++ {
		greetings[i] = &model.Greeting{
			ID:      fmt.Sprintf("stream-%d", i),
			Name:    name,
			Message: "Hello " + name + "!",
		}
	}
	return greetings, nil
}

// GetUserGreetings はユーザーの挨拶一覧を取得（交差テーブル経由）
func (u *greeterUsecase) GetUserGreetings(ctx context.Context, userID string) ([]*model.Greeting, error) {
	return u.userGreetingRepo.FindGreetingsByUserID(ctx, userID)
}
