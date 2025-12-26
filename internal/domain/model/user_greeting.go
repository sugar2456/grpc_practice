package model

import "time"

// UserGreeting はユーザーと挨拶の関連（交差テーブル）
type UserGreeting struct {
	ID         string
	UserID     string
	GreetingID string
	CreatedAt  time.Time
}

// NewUserGreeting は新しいUserGreetingを作成
func NewUserGreeting(id, userID, greetingID string) *UserGreeting {
	return &UserGreeting{
		ID:         id,
		UserID:     userID,
		GreetingID: greetingID,
		CreatedAt:  time.Now(),
	}
}
