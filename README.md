# grpc_practice

Connect-go を使った gRPC サービスのサンプルです。Connect は gRPC、gRPC-Web、Connect プロトコルをサポートし、ブラウザから直接アクセスできます。

---

## 前提（Prerequisites）

- Docker
- Make

---

## セットアップ

```bash
# protoファイルからコードを生成
make proto

# go.sum を更新
make tidy
```

---

## サーバを起動する

```bash
make run
```

サーバがポート `8081` で起動します。

---

## 動作確認

curl で直接 JSON を送れます（Connect の利点）:

```bash
curl -X POST http://localhost:8081/greeter.Greeter/SayHello \
  -H "Content-Type: application/json" \
  -d '{"name": "World"}'
```

期待されるレスポンス:

```json
{"message":"Hello, World!"}
```

---

## Makefile コマンド一覧

| コマンド | 説明 |
|----------|------|
| `make build` | Docker イメージをビルド |
| `make run` | サーバーを起動（ポート 8081） |
| `make shell` | コンテナにシェルで入る |
| `make proto` | proto ファイルをコンパイル |
| `make tidy` | go mod tidy を実行 |
| `make clean` | 生成ファイルを削除 |

ポートを変更したい場合:

```bash
make run PORT=9000
```

---

## プロジェクト構成

```
grpc_practice/
├── Dockerfile
├── Makefile
├── go.mod
├── proto/
│   └── greeter.proto      # サービス定義
├── gen/                   # 生成コード（make proto で生成）
│   └── greeter/v1/
│       ├── greeter.pb.go
│       └── greeterv1connect/
│           └── greeter.connect.go
└── cmd/
    └── server/
        └── main.go        # サーバー実装
```

---

## gRPC vs Connect

| 項目 | gRPC | Connect |
|------|------|---------|
| ブラウザ対応 | Envoy プロキシ必要 | 直接対応 |
| HTTP | HTTP/2 のみ | HTTP/1.1, HTTP/2 |
| 動作確認 | grpcurl 必要 | curl で OK |

---

## 参考

- [Connect-go 公式ドキュメント](https://connectrpc.com/docs/go/getting-started)
- proto: `proto/greeter.proto`
- 生成コード: `gen/greeter/v1/`
- サーバー実装: `cmd/server/main.go`
