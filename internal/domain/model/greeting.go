package model

// Greeting はドメインエンティティ
type Greeting struct {
	ID      string
	Name    string
	Message string
}

// NewGreeting は新しいGreetingを作成
func NewGreeting(id, name string) *Greeting {
	return &Greeting{
		ID:      id,
		Name:    name,
		Message: "Hello, " + name + "!",
	}
}
