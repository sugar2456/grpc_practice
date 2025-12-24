FROM golang:1.25

WORKDIR /app

# gRPC用ツールのインストール
RUN apt update && apt install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# grpcurl をインストール（gRPC サービスの簡易検証に便利）
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# PATHにgo binを追加
ENV PATH="$PATH:/root/go/bin"

COPY . .

# go.modがあれば依存をダウンロード
RUN if [ -f go.mod ]; then go mod download; fi

CMD ["sh"]
