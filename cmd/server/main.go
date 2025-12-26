package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"connectrpc.com/connect"
	greeterv1 "grpc_practice/gen/greeter/v1"
	"grpc_practice/gen/greeter/v1/greeterv1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(
	ctx context.Context,
	req *connect.Request[greeterv1.HelloRequest],
) (*connect.Response[greeterv1.HelloReply], error) {
	log.Printf("Request: %v", req.Msg.Name)
	res := connect.NewResponse(&greeterv1.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}

func (s *GreeterServer) SayHelloStream(
	ctx context.Context,
	req *connect.Request[greeterv1.HelloRequest],
	stream *connect.ServerStream[greeterv1.HelloReply],
) error {
	name := req.Msg.Name
	for i := 0; i < 3; i++ {
		if err := stream.Send(&greeterv1.HelloReply{
			Message: fmt.Sprintf("Hello %s! (%d)", name, i+1),
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	greeter := &GreeterServer{}
	mux := http.NewServeMux()
	path, handler := greeterv1connect.NewGreeterHandler(greeter)
	mux.Handle(path, handler)

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

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Stopping server...")
	server.Close()
}
