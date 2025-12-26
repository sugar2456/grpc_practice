FROM golang:1.25

# ホストユーザーと同じUID/GIDでユーザー作成（ビルド時に指定）
ARG UID=1000
ARG GID=1000
RUN groupadd -g ${GID} appuser && useradd -m -u ${UID} -g ${GID} appuser

WORKDIR /app

# gRPC用ツールのインストール（rootで実行）
RUN apt update && apt install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*

# appuserに切り替え
USER appuser

# Goツールのインストール（Connect用）
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# PATHにgo binを追加
ENV PATH="$PATH:/home/appuser/go/bin"

CMD ["sh"]
