FROM golang:1.22 AS golang
FROM node:20 AS nodejs

# 新しいイメージの作成
FROM ubuntu:22.04

# 開発で使いそうなツールのインストール
RUN apt-get update && \
    # なんか非推奨らしい?
    # apt-get upgrade -y && \
    apt-get install -y \
    # 最小限でいいときは以下にする
    # apt-get install -y --no-install-recommends
    build-essential \
    curl \
    git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Go と Node.js から 動かすために必要な要素を コピー
COPY --from=golang /usr/local/go /usr/local/go
COPY --from=nodejs /usr/local/bin/node /usr/local/bin/
COPY --from=nodejs /usr/local/lib/node_modules /usr/local/lib/node_modules

# Go と Node.js の パスを通す
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/usr/local/bin:${PATH}"

# 生成させたら COPY だったが 動かなかったので コマンド本体 と bin 内のコマンド で エイリアス を作った
# COPY --from=nodejs /usr/local/bin/npm /usr/local/bin/
# COPY --from=nodejs /usr/local/bin/npx /usr/local/bin/
RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm
RUN ln -s /usr/local/lib/node_modules/npm/bin/npx-cli.js /usr/local/bin/npx

# Google CLI のインストール
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
    gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg && \
    apt-get update -y && \
    apt-get install google-cloud-sdk -y

# バックエンド の Go で 必要なライブラリをインストール
WORKDIR /app/backend
COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

# フロントエンド の Node.js で 必要なライブラリをインストール
WORKDIR /app/frontend/local_test_svelte/
COPY ../frontend/local_test_svelte/package*.json ./
RUN npm install

# 本当は ソースのコピーと実行コマンドを定義しないといけない
# この image の使い方が docker compose からマウントして
# bash で 動かす開発環境だから CMD も書かない
