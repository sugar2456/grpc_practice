package model

// Greeting はドメインエンティティ
type Greeting struct {
	Name    string
	Message string
}

// NewGreeting は新しいGreetingを作成
func NewGreeting(name string) *Greeting {
	return &Greeting{
		Name:    name,
		Message: "Hello, " + name + "!",
	}
}
