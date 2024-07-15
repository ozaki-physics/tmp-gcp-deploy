# ビルドステージ
FROM golang:1.22 as builder

# 作業ディレクトリの設定
WORKDIR /test/backend

# Goモジュールをキャッシュ
COPY go.mod .
COPY go.sum .
RUN go mod download

# アプリケーションのソースをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main main.go

# アプリケーションを実行
# CMD ["go", "run", "main.go"]
CMD ["./main"]


# # 実行ステージ
# FROM alpine:3.16.3

# # ビルド済みのバイナリをコピー
# COPY --from=builder /test/backend/main /test/backend/main
