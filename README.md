# grpc_practice

簡単な gRPC サービスのサンプルです（`proto/greeter.proto` → 生成された `pb/*.pb.go` が含まれます）。この README では、開発環境の前提、プロトコルバッファの再生成方法、サーバ起動と動作確認手順を記載しています。 ✅

---

## 前提（Prerequisites） 🔧
- Go 1.24 以上
- protoc（Protocol Buffers コンパイラ）
- 以下のツール（必要に応じてインストール）:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  - `grpcurl`（動作確認用）

インストール例（Go のツール）:

```bash
# protoc-gen-go / protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# grpcurl (コマンドラインから gRPC を叩く場合)
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```


---

## .proto からコードを再生成する
`proto/greeter.proto` を編集した場合、Go 用のコードを再生成してください:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/greeter.proto
```

生成されるファイル: `grpc_practice/pb/greeter.pb.go`, `grpc_practice/pb/greeter_grpc.pb.go`（既に含まれています）

> **注意**: 生成ファイルは直接編集しないでください。`.proto` を編集して再生成するのが正しいワークフローです。

---

## サーバをビルド / 起動する

プロジェクト全体のビルド確認:

```bash
go build ./...
```

サーバを起動する（開発用、ポート `:50051`）:

```bash
# フォアグラウンドで起動 (Ctrl+C で停止)
go run ./cmd/server

# バックグラウンドで起動してログを残す
nohup go run ./cmd/server > server.log 2>&1 &
```

起動するとログに `server started at :50051` が出ます。

---

## grpcurl で動作確認（おすすめ）

`grpcurl` を使うと、ローカルの gRPC サーバを簡単に呼べます（プレーンテキスト接続）:

```bash
grpcurl -plaintext -import-path proto -proto proto/greeter.proto \
  -d '{"name":"Alice"}' localhost:50051 greeter.Greeter/SayHello
```

期待されるレスポンス:

```json
{ "message": "Hello Alice" }
```

---

## サーバ停止方法
- フォアグラウンド: Ctrl+C
- バックグラウンド: プロセスを確認して `kill <PID>` または `pkill -f "go run ./cmd/server"`

---

## 追加案 (任意)
- `bufconn` を使ったユニットテストの追加（ネットワーク不要で CI 向け）
- `grpc-gateway` を導入して `curl` で確認できる HTTP/JSON エンドポイントを用意

---

## 参考
- proto: `proto/greeter.proto`
- 生成コード: `grpc_practice/pb/`
- サーバ実装: `cmd/server/main.go`

---

何か追記してほしい内容（例：Docker 起動手順、Makefile、CI の設定など）があれば教えてください。