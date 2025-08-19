# ベースイメージを指定
FROM golang:latest

# tsコマンドはログのタイムスタンプを付与するために使用
# curlコマンドはAtlas CLIのインストールに使用
RUN apt-get update && apt-get install -y \
    moreutils \
    curl \
    && rm -rf /var/lib/apt/lists/*

# タイムゾーンを設定
ENV TZ=Asia/Tokyo

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# 開発ツールのインストール
RUN go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install entgo.io/ent/cmd/ent@latest

# Atlas CLI をインストール（プリビルトバイナリ）
RUN curl -sSf https://atlasgo.sh | sh -s -- -b /usr/local/bin

# ソースコードをコピー
COPY . .

# Go Modules の依存解決
RUN go mod download

# 公開予定のコンテナのポートを明示
EXPOSE 8080

# Air(ホットリロード)を実行, 標準出力と標準エラー出力をログファイルに出力
CMD ["sh", "-c", "air 2>&1 | ts '[%Y/%m/%d %H:%M:%S]' | tee tmp/build-output.log"]
