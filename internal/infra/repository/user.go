package repository

import (
	"context"
	"sync"

	"grpc_practice/internal/domain/model"
	"grpc_practice/internal/domain/repository"
)

// InMemoryUserRepository はメモリ内ユーザーリポジトリ実装
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*model.User
}

// NewInMemoryUserRepository は新しいInMemoryUserRepositoryを作成
func NewInMemoryUserRepository() repository.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]*model.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}
