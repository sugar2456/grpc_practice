package handler

import (
	"context"
	"log"

	"connectrpc.com/connect"

	greeterv1 "grpc_practice/gen/greeter/v1"
	"grpc_practice/internal/usecase"
)

// GreeterHandler はConnectのハンドラー実装
type GreeterHandler struct {
	usecase usecase.GreeterUsecase
}

// NewGreeterHandler は新しいGreeterHandlerを作成
func NewGreeterHandler(uc usecase.GreeterUsecase) *GreeterHandler {
	return &GreeterHandler{usecase: uc}
}

func (h *GreeterHandler) SayHello(
	ctx context.Context,
	req *connect.Request[greeterv1.HelloRequest],
) (*connect.Response[greeterv1.HelloReply], error) {
	log.Printf("Request: %v", req.Msg.Name)

	// Usecaseを呼び出し
	greeting, err := h.usecase.SayHello(ctx, req.Msg.Name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスに変換
	res := connect.NewResponse(&greeterv1.HelloReply{
		Message: greeting.Message,
	})
	return res, nil
}

func (h *GreeterHandler) SayHelloStream(
	ctx context.Context,
	req *connect.Request[greeterv1.HelloRequest],
	stream *connect.ServerStream[greeterv1.HelloReply],
) error {
	log.Printf("Stream Request: %v", req.Msg.Name)

	// Usecaseを呼び出し
	greetings, err := h.usecase.SayHelloStream(ctx, req.Msg.Name, 3)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}

	// ストリームでレスポンスを送信
	for i, g := range greetings {
		if err := stream.Send(&greeterv1.HelloReply{
			Message: g.Message + " (" + string(rune('1'+i)) + ")",
		}); err != nil {
			return err
		}
	}
	return nil
}
