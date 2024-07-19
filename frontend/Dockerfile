# 1. ビルド環境
FROM node:20 AS builder

WORKDIR /test/frontend/local_test_svelte

# パッケージをコピーしてインストール
COPY ./local_test_svelte/package*.json ./
RUN npm install

# アプリケーションのソースをコピー
COPY ./local_test_svelte .
# COPY ./local_test_svelte/src ./src
# COPY ./local_test_svelte/static ./static
# COPY ./local_test_svelte/svelte.config.js ./svelte.config.js
# COPY ./local_test_svelte/vite.config.ts ./vite.config.ts

# ビルドする
RUN npm run build

# 2. 実行環境
FROM node:20-slim AS stage-build

WORKDIR /app

# パッケージをコピーしてインストール
COPY ./local_test_svelte/package*.json ./
RUN npm ci --omit dev

# Copy the built project files from the builder stage
COPY --from=builder /test/frontend/local_test_svelte/build /app/build

# Set environment variable to indicate the production environment
ENV NODE_ENV=production

# Command to run the application
CMD ["node", "build"]
