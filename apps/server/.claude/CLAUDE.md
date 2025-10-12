# Recall GraphQL Server

## 技術スタック

### 言語・フレームワーク
- **Go** 1.24.0
- **GraphQL**: gqlgen
- **HTTP Server**: Echo
- **ORM**: GORM
- **Migration**: Goose

### 認証・セキュリティ
- **JWT**: golang-jwt
- **認証**: Supabase Auth (SSO専用)

### データベース・ストレージ
- **Database**: PostgreSQL (Supabase)
- **Object Storage**: S3 (MinIO)

### インフラ・運用
- **Deploy**: AWS Lambda
- **Monitoring**: Sentry
- **Hot Reload**: Air

## アーキテクチャ

- **クリーンアーキテクチャ**を採用
- 参考記事: https://zenn.dev/88888888_kota/articles/1a5caac8d743b8

### ディレクトリ構成
```
apps/server/
├── cmd/           # エントリーポイント
├── config/        # 設定ファイル
├── graph/         # GraphQLスキーマ・リゾルバー
├── migrations/    # データベースマイグレーション
├── pkg/
│   ├── adapter/   # 外部アダプター
│   │   ├── controller/
│   │   └── middleware/
│   ├── domain/    # ドメイン層
│   │   ├── entity/
│   │   └── repository/
│   ├── infra/     # インフラ層
│   │   ├── database/
│   │   └── storage/
│   └── usecase/   # ユースケース層
└── server.go      # サーバー起動
```

## 設計指針

クリーンアーキテクチャに従った層分離により、テスタブルで保守性の高いコードベースを目指しています。