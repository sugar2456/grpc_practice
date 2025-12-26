.PHONY: build run shell proto tidy clean

# Dockerイメージ名
IMAGE_NAME := grpc-go
# ポート番号（デフォルト: 8081）
PORT := 8081

# Dockerイメージをビルド
build:
	docker build -t $(IMAGE_NAME) .

# コンテナ内でサーバーを起動
run: build
	docker run -it --rm -v $(PWD):/app -p $(PORT):8080 $(IMAGE_NAME) go run cmd/server/main.go

# コンテナにシェルで入る
shell: build
	docker run -it --rm -v $(PWD):/app -p $(PORT):8080 $(IMAGE_NAME)

# protoファイルをコンパイル
proto: build
	docker run --rm -v $(PWD):/app $(IMAGE_NAME) protoc --go_out=module=grpc_practice:. --connect-go_out=module=grpc_practice:. proto/greeter.proto

# go mod tidy
tidy: build
	docker run --rm -v $(PWD):/app $(IMAGE_NAME) go mod tidy

# 生成ファイルを削除
clean:
	rm -rf gen/
