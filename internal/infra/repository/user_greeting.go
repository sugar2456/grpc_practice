package repository

import (
	"context"
	"sync"

	"grpc_practice/internal/domain/model"
	"grpc_practice/internal/domain/repository"
)

// InMemoryUserGreetingRepository は交差テーブルのメモリ内実装
type InMemoryUserGreetingRepository struct {
	mu            sync.RWMutex
	userGreetings map[string]*model.UserGreeting
	// 他のリポジトリへの参照（JOINに使用）
	userRepo     repository.UserRepository
	greetingRepo repository.GreetingRepository
}

// NewInMemoryUserGreetingRepository は新しいInMemoryUserGreetingRepositoryを作成
func NewInMemoryUserGreetingRepository(
	userRepo repository.UserRepository,
	greetingRepo repository.GreetingRepository,
) repository.UserGreetingRepository {
	return &InMemoryUserGreetingRepository{
		userGreetings: make(map[string]*model.UserGreeting),
		userRepo:      userRepo,
		greetingRepo:  greetingRepo,
	}
}

func (r *InMemoryUserGreetingRepository) Save(ctx context.Context, ug *model.UserGreeting) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userGreetings[ug.ID] = ug
	return nil
}

func (r *InMemoryUserGreetingRepository) FindByUserID(ctx context.Context, userID string) ([]*model.UserGreeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.UserGreeting, 0)
	for _, ug := range r.userGreetings {
		if ug.UserID == userID {
			result = append(result, ug)
		}
	}
	return result, nil
}

func (r *InMemoryUserGreetingRepository) FindByGreetingID(ctx context.Context, greetingID string) ([]*model.UserGreeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.UserGreeting, 0)
	for _, ug := range r.userGreetings {
		if ug.GreetingID == greetingID {
			result = append(result, ug)
		}
	}
	return result, nil
}

// FindGreetingsByUserID はユーザーIDで挨拶一覧を取得（JOINイメージ）
func (r *InMemoryUserGreetingRepository) FindGreetingsByUserID(ctx context.Context, userID string) ([]*model.Greeting, error) {
	userGreetings, err := r.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	greetings := make([]*model.Greeting, 0, len(userGreetings))
	for _, ug := range userGreetings {
		greeting, err := r.greetingRepo.FindByID(ctx, ug.GreetingID)
		if err != nil {
			return nil, err
		}
		if greeting != nil {
			greetings = append(greetings, greeting)
		}
	}
	return greetings, nil
}

// FindUsersByGreetingID は挨拶IDでユーザー一覧を取得（JOINイメージ）
func (r *InMemoryUserGreetingRepository) FindUsersByGreetingID(ctx context.Context, greetingID string) ([]*model.User, error) {
	userGreetings, err := r.FindByGreetingID(ctx, greetingID)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0, len(userGreetings))
	for _, ug := range userGreetings {
		user, err := r.userRepo.FindByID(ctx, ug.UserID)
		if err != nil {
			return nil, err
		}
		if user != nil {
			users = append(users, user)
		}
	}
	return users, nil
}
