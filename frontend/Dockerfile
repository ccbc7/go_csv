FROM node:current-slim

WORKDIR /app

# パッケージファイルをコピー
COPY package*.json ./

# 依存関係をインストール（キャッシュをクリアして完全にインストール）
RUN npm cache clean --force && npm install --include=dev

# アプリケーションファイルをコピー
COPY . .

# 開発環境用：ファイル権限を777に設定
RUN chmod -R 777 /app

# ポート3000はアプリケーションのポート、9229はデバッグポート
EXPOSE 3000 9229

# 環境変数で開発/本番を切り替え
ENV NODE_ENV=development

# 開発サーバーを起動（デバッグモードで実行）
CMD ["node", "--inspect=0.0.0.0:9229", "./node_modules/@remix-run/dev/dist/cli.js", "vite:dev", "--host", "--port", "3000"]

ENTRYPOINT ["/entrypoint.sh"]
