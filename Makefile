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
migrate:
	docker-compose run --rm backend go run cmd/project/main.go -migrate

seed:
	docker-compose run --rm backend go run cmd/project/main.go -seed

# マイグレーションとシードを順次実行
setup: migrate seed

# データベースを完全リセット
reset:
	docker-compose run --rm backend go run cmd/project/main.go -reset

# シーケンスをリセット（IDを1から開始）- 全テーブルのシーケンスを自動検出
reset_sequences:
	docker-compose exec db psql -U ginuser -d gin -c "SELECT setval(sequence_name::text, 1, false) FROM information_schema.sequences WHERE sequence_schema = 'public' AND sequence_name LIKE '%_id_seq';"

# データベースを完全リセットしてシーケンスもリセット、シードも実行
reset_setup: 
	$(MAKE) reset
	$(MAKE) reset_sequences  
	$(MAKE) seed

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
	docker-compose run --rm backend go generate ./internal/ent

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
