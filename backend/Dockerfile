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

# アプリケーションを実行
# CMD ["./main"]
RUN ls -al


# 実行ステージ
FROM alpine:3

# ビルド済みのバイナリをコピー
COPY --from=builder /test/backend/main main
RUN ls -al

CMD ["./main"]
