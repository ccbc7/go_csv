FROM node:current-slim

WORKDIR /app

# パッケージファイルをコピー
COPY package*.json ./

# 依存関係をインストール（キャッシュをクリアして完全にインストール）
RUN npm cache clean --force && npm install --include=dev

# アプリケーションファイルをコピー
COPY . .

# ファイルの所有者をnodeユーザーに変更
RUN chown -R node:node /app

# nodeユーザーに切り替え
USER node

# ポート3000と9229を公開
EXPOSE 3000 9229

# 環境変数で開発/本番を切り替え
ENV NODE_ENV=development

# 開発サーバーを起動（デバッグモードで実行）
CMD ["node", "--inspect=0.0.0.0:9229", "./node_modules/@remix-run/dev/dist/cli.js", "vite:dev", "--host", "--port", "3000"]
