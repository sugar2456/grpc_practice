package main

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"

	"grpc_practice/gen/greeter/v1/greeterv1connect"
	"grpc_practice/internal/domain/repository"
	"grpc_practice/internal/handler"
	infraRepo "grpc_practice/internal/infra/repository"
	"grpc_practice/internal/usecase"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	fx.New(
		// 依存性の提供
		fx.Provide(
			// Infra層
			fx.Annotate(
				infraRepo.NewInMemoryGreetingRepository,
				fx.As(new(repository.GreetingRepository)),
			),
			// Usecase層
			usecase.NewGreeterUsecase,
			// Handler層
			handler.NewGreeterHandler,
			// HTTPサーバー
			NewHTTPServer,
		),
		// サーバー起動
		fx.Invoke(RegisterRoutes),
	).Run()
}

// NewHTTPServer はHTTPサーバーを作成
func NewHTTPServer(lc fx.Lifecycle) *http.ServeMux {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server started at :8080")
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Printf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// RegisterRoutes はルーティングを登録
func RegisterRoutes(mux *http.ServeMux, h *handler.GreeterHandler) {
	path, connectHandler := greeterv1connect.NewGreeterHandler(h)
	mux.Handle(path, connectHandler)
}
