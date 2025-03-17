# ビルド
build:
	docker compose build

# 起動
up:
	docker compose up
upd:
	docker compose up -d

# 再起動
re:
	docker compose restart

# 停止
down:
	docker compose down --remove-orphans

# コンテナ内に入る
b:
	docker compose exec backend bash
f:
	docker compose exec frontend sh


# マイグレーション＆シード
seed:
	docker-compose run --rm backend go run database/migrations/migration.go

# swaggoによるAPIドキュメントの生成&整形
api:
	docker-compose run --rm backend swag init -g cmd/project/main.go && docker-compose run --rm backend swag fmt

# フォーマット（標準のgo fmtを使用）
fmt:
	docker-compose run --rm backend go fmt ./...

# 静的解析（標準のgo vetを使用）
vet:
	docker-compose run --rm backend go vet ./...

# フォーマットと静的解析を実行
fix: fmt vet
