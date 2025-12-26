package repository

import (
	"context"
	"sync"

	"grpc_practice/internal/domain/model"
	"grpc_practice/internal/domain/repository"
)

// InMemoryGreetingRepository はメモリ内リポジトリ実装
type InMemoryGreetingRepository struct {
	mu        sync.RWMutex
	greetings map[string]*model.Greeting // ID -> Greeting
	byName    map[string]*model.Greeting // Name -> Greeting
}

// NewInMemoryGreetingRepository は新しいInMemoryGreetingRepositoryを作成
func NewInMemoryGreetingRepository() repository.GreetingRepository {
	return &InMemoryGreetingRepository{
		greetings: make(map[string]*model.Greeting),
		byName:    make(map[string]*model.Greeting),
	}
}

func (r *InMemoryGreetingRepository) Save(ctx context.Context, greeting *model.Greeting) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.greetings[greeting.ID] = greeting
	r.byName[greeting.Name] = greeting
	return nil
}

func (r *InMemoryGreetingRepository) FindByID(ctx context.Context, id string) (*model.Greeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if g, ok := r.greetings[id]; ok {
		return g, nil
	}
	return nil, nil
}

func (r *InMemoryGreetingRepository) FindByName(ctx context.Context, name string) (*model.Greeting, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if g, ok := r.byName[name]; ok {
		return g, nil
	}
	return nil, nil
}
