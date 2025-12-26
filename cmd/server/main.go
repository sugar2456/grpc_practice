package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"grpc_practice/gen/greeter/v1/greeterv1connect"
	"grpc_practice/internal/handler"
	infraRepo "grpc_practice/internal/infra/repository"
	"grpc_practice/internal/usecase"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// DI: 依存性の注入
	// Infra層（リポジトリ実装）
	greetingRepo := infraRepo.NewInMemoryGreetingRepository()

	// Usecase層
	greeterUsecase := usecase.NewGreeterUsecase(greetingRepo)

	// Handler層
	greeterHandler := handler.NewGreeterHandler(greeterUsecase)

	// ルーティング
	mux := http.NewServeMux()
	path, h := greeterv1connect.NewGreeterHandler(greeterHandler)
	mux.Handle(path, h)

	// サーバー起動
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		log.Println("Server started at :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Stopping server...")
	server.Close()
}
