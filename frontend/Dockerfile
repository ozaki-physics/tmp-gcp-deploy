# 1. ビルド環境
FROM node:20 AS builder

WORKDIR /test/frontend/local_test_svelte

# パッケージをコピーしてインストール
COPY ./local_test_svelte/package*.json ./
RUN npm install

# アプリケーションのソースをコピー
COPY ./local_test_svelte .

# ビルドする
RUN npm run build

# ビルド出力ディレクトリを確認
RUN ls -al /test/frontend/local_test_svelte/build

# 2. 実行環境
FROM node:20 AS stage-build

WORKDIR /app

# パッケージをコピーしてインストール
COPY ./local_test_svelte/package*.json ./
RUN npm ci --omit dev

# Copy the built project files from the builder stage
COPY --from=builder /test/frontend/local_test_svelte /app/build

# Set environment variable to indicate the production environment
ENV NODE_ENV=production

# Command to run the application
CMD ["node", "build"]
