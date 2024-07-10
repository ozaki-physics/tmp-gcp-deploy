# ベースイメージ
FROM node:20

# 作業ディレクトリの設定
WORKDIR /test/frontend/local_test_svelte

# # パッケージをコピーしてインストール
# COPY ./local_test_svelte/package*.json .
# RUN npm install

# # アプリケーションのソースをコピー
# COPY ./local_test_svelte .

# 開発サーバーを起動
CMD ["npm", "run", "dev", "--", "--host", "--port=8080"]
