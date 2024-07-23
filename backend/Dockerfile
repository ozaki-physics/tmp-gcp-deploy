# ビルドステージ
FROM golang:1.22 AS builder

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

# 実行ステージ
FROM debian:stable-slim

WORKDIR /app

# なぜか path を通さないとエラーになった
ENV PATH="/app:${PATH}"

# ビルド済みのバイナリをコピー
COPY --from=builder /test/backend/main /app/main

CMD ["./main"]
