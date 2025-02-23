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
	docker compose exec frontend ash
