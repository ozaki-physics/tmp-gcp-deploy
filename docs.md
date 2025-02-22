### 使い方
#### 全部入りの 1個のコンテナ で動かすとき
カレントディレクトリ: `tmp-gcp-deploy`
コマンド: 
`docker-compose -f .devcontainer/docker-compose.yml build`

ビルドするのに 37.2 sec
でも Node で npm をインストールしたりしてない時間
`docker-compose -f .devcontainer/docker-compose.yml up -d`
`docker-compose -f .devcontainer/docker-compose.yml exec local_test bash`
`docker-compose -f .devcontainer/docker-compose.yml down`
たしかに 使えるようになったっぽい
```bash
$ gcloud --version
Google Cloud SDK 480.0.0
alpha 2024.06.07
beta 2024.06.07
bq 2.1.5
bundled-python3-unix 3.11.8
core 2024.06.07
gcloud-crc32c 1.0.0
gsutil 5.29
```
#### フロント と バック の コンテナ をそれぞれ 動かすとき
カレントディレクトリ: `tmp-gcp-deploy`
`docker-compose build`

#### メモ
なんか
npm install --production じゃなくて
npm install --omit=dev が推奨らしい
https://zenn.dev/zawa_kyo/articles/d671e0935ae0c0

どうやら ちゃんと 
`npm i -D @sveltejs/adapter-node` して
svelte.config.js に
`import adapter from '@sveltejs/adapter-node';` を書かないと
`npm run build` しても build ディレクトリが生成されなかった
https://kit.svelte.jp/docs/adapter-node

package.json の scripts は エイリアス な感じで npm run XXX か npm XXX で動くらしい
だから `npm run build` しても 実際に動くのは `vite build` だった
ここに `"start": "node build",` を追加したけど 必ず必要なわけではない

コンテナ動かそうとして 以下のエラーに悩まされた
```
------
 > [frontend internal] load build context:
------
failed to solve: archive/tar: unknown file mode ?rwxr-xr-x
```
結論はよく分からなかったけど おそらく COPY しようとしてるファイルが多すぎるからかと
GitHub Actions の場合は そもそも push してないから 問題にならなかった
以下にしたら image 作れた
```dockerfile
# COPY ./local_test_svelte .
COPY ./local_test_svelte/src ./src
COPY ./local_test_svelte/static ./static
COPY ./local_test_svelte/svelte.config.js ./svelte.config.js
COPY ./local_test_svelte/vite.config.ts ./vite.config.ts
```

- Docker のディスク使用量を表示
https://matsuand.github.io/docs.docker.jp.onthefly/engine/reference/commandline/system_df/
`docker system df`
- 構築キャッシュを削除
https://matsuand.github.io/docs.docker.jp.onthefly/engine/reference/commandline/builder_prune/
`docker builder prune`

.devcontainer/docker-compose.yml で
ホストの node_modules をマウントしたくないからと volume を書いたが docker-compose 内で volume を明記しないと volume が量産されてしまった
`docker volume ls`

悩まされてたエラー
最初は main が見当たらない だった
>terminated: Application failed to start: failed to load /main: no such file or directory
途中から 以下の 実行できないになった root ユーザーとかも確認したけど パスを通す以外の方法が分からなかった
>Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "main": executable file not found in $PATH: unknown

なんで どのドキュメントにも EMV でパスを通すことは書かれてないのに そんなエラーに見舞われたんだろう

なんか今は Distroless という軽量 image があるらしい
ドキュメント: 軽量Dockerイメージに安易にAlpineを使うのはやめたほうがいいという話
https://blog.inductor.me/entry/alpine-not-recommended

`docker-compose build frontend --progress=plain` を実行してたら
>--progress is a global compose flag, better use `docker compose --progress xx build ...
と言われた

docker-compose で 操作する サービス を指定する(frontend とか backend とか)

- backend のとき
docker-compose build backend --progress=plain
docker-compose up -d backend
docker container ls -a
docker-compose down backend
docker image ls -a
docker image rm local_go:test
docker builder prune

docker-compose exec backend bash
docker container run -it --rm -p 5001:8081 local_go:test bash


- frontend のとき
(
  docker compose と ハイフンが必要なくなったらしい
  今まで docker compose が別アプリだったから 1個のコマンドとしないといけなかった? 最近は docker に組み込まれたから 必要ない?
)
docker compose --progress plain build frontend
docker compose up -d frontend
docker container ls -a
docker compose down frontend
docker image ls -a
docker image rm local_svelte:test
docker builder prune

docker compose exec frontend bash
docker container run -it --rm -p 5002:8081 local_go:test bash

`failed to solve: archive/tar: unknown file mode ?rwxr-xr-x` が解決できない

node_modules を 削除したら 正常に動いた
.svelte-kit は残ったままでも大丈夫
つまり Linux で動かしてるのに Windows にマウントしたときに 壊れちゃうっぽい

>npm warn deprecated inflight@1.0.6: This module is not supported, and leaks memory. Do not use it. Check out lru-cache if you want a good and tested way to coalesce async requests by a key value, which is much more comprehensive and powerful.
おい メモリリークするような ライブラリがあるんだが?

npm run dev
