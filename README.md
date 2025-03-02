# gin_csv

## バックエンド

### アプリ

http://localhost:8080

### swagger

http://localhost:8080/swagger/index.html

### ディレクトリ構成

（参考）https://github.com/golang-standards/project-layout/blob/master/README_ja.md

```
　/project-root
│── /cmd              # エントリーポイント
│── /config           # 設定ファイル
│── /internal　(外部からアクセスしないパッケージ)
│   │── /handler      # リクエスト処理(コントローラー)
│   │── /service      # ビジネスロジック（FATコントローラーを防ぐ）
│   │── /repository   # データベース操作
│   │── /request      # リクエストのバリデーション
│   │── /dto          # レスポンスDTO（データ転送オブジェクト）
│   │── /model        # エンティティ
│   │── /middleware   # 認証・リクエスト処理前のロジック
│   │── /routes       # ルーティング定義
│── /pkg              # 汎用ライブラリ
│── main.go
```

### フレームワーク

- [gin](https://gin-gonic.com/ja/docs/)
  <br> Go 言語向けの軽量かつ高性能な HTTPWeb フレームワーク

### ライブラリ

- [gorm](https://gorm.io/)
  <br> ORM。SQL を書かなくてもデータベース操作ができる
- [gormigrate](https://github.com/go-gormigrate/gormigrate)
  <br> gorm だけでは実現できないマイグレーションのバージョン管理ができる
- [swaggo](https://github.com/swaggo/swag)
  <br> swagger ドキュメントの自動生成と swagger-ui による画面上での API 実行の機能の提供
- [air](https://github.com/air-verse/air)
  <br> ホットリロードを提供するツール
- [delve](https://github.com/go-delve/delve)
  <br> デバッガー。Xdebug のような使い方ができる
