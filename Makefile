# ビルド
build:
	docker compose build

# ビルド（キャッシュなし）
build_no_cache:
	docker compose build --no-cache

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

# 依存関係の解決
tidy:
	docker-compose run --rm backend go mod tidy

# 型安全なORM用のメソッドの生成
ent:
	docker-compose run --rm backend ent generate --target ./internal/ent ./schema

# Atlasのマイグレーションの差分を生成　make atlas_diff xxx
atlas_diff:
	docker-compose run --rm backend atlas migrate diff $(filter-out $@,$(MAKECMDGOALS)) \
--dir "file:///./internal/database/migrations" \
--to "ent://schema" \
--dev-url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"

# Atlasマイグレーション削除(※手動で削除してもvolumeを消さないとAtlasは気づかないよ)
atlas_rm:
	docker-compose run --rm backend atlas migrate rm \
--dir "file://internal/database/migrations"

# # Atlasマイグレーション適用
# atlas_apply:
# 	docker-compose run --rm backend atlas migrate apply \
# --dir "file://internal/database/migrations" \
# --url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"

# # Atlasマイグレーション状態確認
# atlas_status:
# 	docker-compose run --rm backend atlas migrate status \
# --dir "file://internal/database/migrations" \
# --url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"

# 位置引数を無視するためのダミーターゲット
%:
	@:
