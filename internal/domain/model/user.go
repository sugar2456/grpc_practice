package model

// User はユーザーエンティティ
type User struct {
	ID   string
	Name string
}

// NewUser は新しいUserを作成
func NewUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
