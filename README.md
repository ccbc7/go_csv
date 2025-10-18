# gin_csv

Go 言語の Gin フレームワークを使用した REST API アプリケーションです。アイテム管理と CSV 操作機能を提供し、JWT 認証によるセキュアな API アクセスを実現しています。

## 🚀 クイックスタート

### 前提条件

- Docker & Docker Compose
- Make（オプション）

### セットアップ

1. リポジトリをクローン

```bash
git clone <repository-url>
cd gin_csv
```

2. 環境変数の設定

```bash
cp .env.example .env  # 必要に応じて値を調整
```

3. アプリケーションの起動

```bash
make up
# または
docker compose up -d
```

4. データベースの初期化（初回のみ）

```bash
make seed
```

### アクセス情報

- **アプリケーション**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **PostgreSQL**: localhost:5432

## 📁 プロジェクト構成

```
gin_csv/
├── backend/                    # バックエンドアプリケーション
│   ├── cmd/project/           # アプリケーションエントリーポイント
│   │   └── main.go           # メイン実行ファイル
│   ├── config/               # 設定管理
│   │   └── env.go           # 環境変数設定
│   ├── database/            # データベース関連
│   │   ├── db_setup.go     # DB接続設定
│   │   ├── migrations/     # マイグレーション
│   │   └── seeders/        # シードデータ
│   ├── internal/           # 内部パッケージ
│   │   ├── dto/           # データ転送オブジェクト
│   │   ├── handlers/      # HTTPハンドラー（コントローラー）
│   │   ├── middlewares/   # ミドルウェア
│   │   ├── models/        # データモデル
│   │   ├── repositories/  # データアクセス層
│   │   ├── router/        # ルーティング設定
│   │   ├── services/      # ビジネスロジック
│   │   ├── scripts/       # ユーティリティスクリプト
│   │   ├── tests/         # テストファイル
│   │   └── utils/         # 共通ユーティリティ
│   ├── docs/              # API仕様書（Swagger）
│   ├── tmp/               # 一時ファイル（Air用）
│   ├── .air.toml          # Air設定（ホットリロード）
│   ├── Dockerfile         # Dockerイメージ定義
│   ├── go.mod             # Go依存関係
│   └── go.sum             # Go依存関係チェックサム
├── compose.yaml            # Docker Compose設定
├── Makefile               # 開発用コマンド
└── README.md              # このファイル
```

## 🛠 技術スタック

### フレームワーク・ライブラリ

- **[Gin](https://gin-gonic.com/ja/docs/)** - 軽量・高性能な HTTP Web フレームワーク
- **[GORM](https://gorm.io/)** - Go 用 ORM（Object-Relational Mapping）
- **[Gormigrate](https://github.com/go-gormigrate/gormigrate)** - マイグレーションバージョン管理
- **[Swaggo](https://github.com/swaggo/swag)** - Swagger API 仕様書自動生成
- **[JWT-Go](https://github.com/golang-jwt/jwt)** - JWT 認証実装
- **[Air](https://github.com/air-verse/air)** - ホットリロード開発ツール
- **[Delve](https://github.com/go-delve/delve)** - Go デバッガー

### データベース

- **PostgreSQL 16** - メインデータベース
- **SQLite** - テスト用データベース（オプション）

### 開発ツール

- **Docker & Docker Compose** - コンテナ化
- **Make** - タスクランナー
- **Air** - ホットリロード


## 🔧 開発コマンド

### Docker 操作

```bash
# アプリケーション起動
make up          # フォアグラウンド実行
make upd         # バックグラウンド実行

# アプリケーション停止
make down

# 再起動
make re

# ビルド
make build

# コンテナ内に入る
make b           # バックエンドコンテナ
```

### データベース操作

```bash
# マイグレーション & シード実行
make seed
```

### API 開発

```bash
# Swagger仕様書生成・整形
make api

# コードフォーマット
make fmt

# 静的解析
make vet

# フォーマット + 静的解析
make fix
```

## 📊 API 仕様

### 認証エンドポイント

| メソッド | エンドポイント | 説明         |
| -------- | -------------- | ------------ |
| POST     | `/auth/signup` | ユーザー登録 |
| POST     | `/auth/login`  | ログイン     |

### アイテム管理エンドポイント

| メソッド | エンドポイント       | 説明             | 認証 |
| -------- | -------------------- | ---------------- | ---- |
| GET      | `/api/v1/items`      | アイテム一覧取得 | 不要 |
| GET      | `/api/v1/items/{id}` | アイテム詳細取得 | 不要 |
| POST     | `/api/v1/items`      | アイテム作成     | 必要 |
| PUT      | `/api/v1/items/{id}` | アイテム更新     | 必要 |
| DELETE   | `/api/v1/items/{id}` | アイテム削除     | 必要 |

### その他

| メソッド | エンドポイント | 説明           |
| -------- | -------------- | -------------- |
| GET      | `/hello`       | ヘルスチェック |

詳細な API 仕様は [Swagger UI](http://localhost:8080/swagger/index.html) で確認できます。

## 🏗 アーキテクチャ

### レイヤー構成

```
┌─────────────────┐
│   Handlers      │ ← HTTPリクエスト処理
│  (Controllers)  │
├─────────────────┤
│    Services     │ ← ビジネスロジック
├─────────────────┤
│  Repositories   │ ← データアクセス
├─────────────────┤
│     Models      │ ← データモデル
└─────────────────┘
```

### 主要コンポーネント

- **Handlers**: HTTP リクエストの受信・レスポンス処理
- **Services**: ビジネスロジックの実装
- **Repositories**: データベース操作の抽象化
- **Models**: データ構造の定義
- **DTOs**: API 入出力データの定義
- **Middlewares**: 認証・CORS・ログ等の横断的関心事

## 🔐 認証・認可

- **JWT（JSON Web Token）**を使用した認証システム
- ログイン時にトークンを発行
- 保護されたエンドポイントでは Authorization ヘッダーでトークンを送信
- トークン形式: `Bearer <token>`

## 🧪 テスト

```bash
# テスト実行
docker compose exec backend go test ./...

# カバレッジ付きテスト実行
docker compose exec backend go test -cover ./...
```

## 📝 開発ガイドライン

### コーディング規約

- Go 標準のフォーマット（`go fmt`）に従う
- `go vet`による静的解析をパスする
- 適切なエラーハンドリングを実装する
- テストコードを記述する

### Git 運用

- feature/機能名 でブランチを作成
- コミット前に `make fix` でコード品質をチェック
- プルリクエスト作成前にテストを実行

## 🚨 トラブルシューティング

### よくある問題

1. **ポート競合エラー**

   ```bash
   # 使用中のポートを確認
   lsof -i :8080
   lsof -i :5432
   ```

2. **データベース接続エラー**

   ```bash
   # コンテナの状態確認
   docker compose ps

   # ログ確認
   docker compose logs db
   ```

3. **マイグレーションエラー**
   ```bash
   # データベースリセット
   make down
   docker volume rm gin_csv_postgres-data
   make up
   make seed
   ```

## 📚 参考資料

- [Go 言語公式ドキュメント](https://golang.org/doc/)
- [Gin 公式ドキュメント](https://gin-gonic.com/docs/)
- [GORM 公式ドキュメント](https://gorm.io/docs/)
- [Go 標準プロジェクトレイアウト](https://github.com/golang-standards/project-layout/blob/master/README_ja.md)

## 📄 ライセンス

このプロジェクトは MIT ライセンスの下で公開されています。
