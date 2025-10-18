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

upf:
	docker compose up -d frontend

upb:
	docker compose up -d backend


# デバッグモードで起動
debug:
	docker compose up -d db
	docker compose run --rm -p 8080:8080 -p 2345:2345 backend dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient cmd/project/main.go


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
migrate:
	docker-compose run --rm backend go run cmd/project/main.go -migrate

seed:
	docker-compose run --rm backend go run cmd/project/main.go -seed

# マイグレーションとシードを順次実行
setup: migrate seed

# データベースを完全リセット
reset:
	docker-compose run --rm backend go run cmd/project/main.go -reset

# swaggoによるAPIドキュメントの生成&整形
swagger:
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
	docker-compose run --rm backend ent generate --target ./internal/ent ./internal/models

# Atlasのマイグレーションファイル生成　make atlas_migrate_diff xxx
atlas_migrate_diff:
	docker compose exec backend atlas migrate diff $(filter-out $@,$(MAKECMDGOALS)) \
--dir "file://internal/database/migrations" \
--to "ent://internal/models" \
--dev-url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"

# Atlasマイグレーションファイルを削除,引数はマイグレーションファイル名(引数はタイムスタンプ 20250921081402のように指定)
atlas_migrate_rm:
	docker compose exec backend atlas migrate rm \
--dir "file://internal/database/migrations" \
$(filter-out $@,$(MAKECMDGOALS))


# Atlasマイグレーション適用
atlas_migrate_apply:
	docker compose exec backend atlas migrate apply \
--dir "file://internal/database/migrations" \
--url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"

# Atlasマイグレーション状態確認
atlas_migrate_status:
	docker compose exec backend atlas migrate status \
--dir "file://internal/database/migrations" \
--url "postgres://ginuser:ginpassword@postgres:5432/gin?sslmode=disable"



#---------------devcontainer---------------
.PHONY: d-f d-b

# frontend起動
d-f:
	npm --prefix ./frontend run dev -- --host --port 3000

# backend起動
d-b:
	cd backend && go run cmd/project/main.go

# 位置引数を無視するためのダミーターゲット（必ず最後）
%:
	@:
