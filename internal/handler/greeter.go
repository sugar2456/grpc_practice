package handler

import (
	"context"
	"fmt"
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
	log.Printf("Request: user_id=%v, name=%v", req.Msg.UserId, req.Msg.Name)

	// Usecaseを呼び出し
	greeting, err := h.usecase.SayHello(ctx, req.Msg.UserId, req.Msg.Name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスに変換
	res := connect.NewResponse(&greeterv1.HelloReply{
		Id:      greeting.ID,
		Message: greeting.Message,
	})
	return res, nil
}

func (h *GreeterHandler) SayHelloStream(
	ctx context.Context,
	req *connect.Request[greeterv1.HelloRequest],
	stream *connect.ServerStream[greeterv1.HelloReply],
) error {
	log.Printf("Stream Request: user_id=%v, name=%v", req.Msg.UserId, req.Msg.Name)

	// Usecaseを呼び出し
	greetings, err := h.usecase.SayHelloStream(ctx, req.Msg.Name, 3)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}

	// ストリームでレスポンスを送信
	for i, g := range greetings {
		if err := stream.Send(&greeterv1.HelloReply{
			Id:      g.ID,
			Message: fmt.Sprintf("%s (%d)", g.Message, i+1),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (h *GreeterHandler) GetUserGreetings(
	ctx context.Context,
	req *connect.Request[greeterv1.GetUserGreetingsRequest],
) (*connect.Response[greeterv1.GetUserGreetingsReply], error) {
	log.Printf("GetUserGreetings: user_id=%v", req.Msg.UserId)

	// Usecaseを呼び出し
	greetings, err := h.usecase.GetUserGreetings(ctx, req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// レスポンスに変換
	protoGreetings := make([]*greeterv1.Greeting, len(greetings))
	for i, g := range greetings {
		protoGreetings[i] = &greeterv1.Greeting{
			Id:      g.ID,
			Name:    g.Name,
			Message: g.Message,
		}
	}

	res := connect.NewResponse(&greeterv1.GetUserGreetingsReply{
		Greetings: protoGreetings,
	})
	return res, nil
}
