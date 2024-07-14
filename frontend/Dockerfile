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

# 2. 実行環境
FROM node:20

# Set working directory
WORKDIR /app

# パッケージをコピーしてインストール
COPY ./local_test_svelte/package*.json ./
RUN npm install --production

# Copy the built project files from the builder stage
COPY --from=builder /app/build /app/build

# Set environment variable to indicate the production environment
ENV NODE_ENV=production

# Command to run the application
CMD ["node", "build"]
