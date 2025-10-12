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

- **GraphQL特化のクリーンアーキテクチャ**を採用
- スキーマ駆動開発（Schema-Driven Development）
- 参考記事: https://qiita.com/WebEngrChild/items/d9b87944235c5220ae5b

### ディレクトリ構成
```
apps/server/
├── cmd/               # エントリーポイント
│   ├── dev/          # 開発用サーバー
│   └── lambda/       # Lambda用ハンドラー
├── migrations/        # データベースマイグレーション
├── pkg/
│   ├── adapter/      # アダプター層
│   │   ├── http/     # HTTP関連
│   │   │   └── resolver/     # GraphQLリゾルバー実装
│   │   │       ├── mutation.resolvers.go
│   │   │       ├── query.resolvers.go
│   │   │       └── resolver.go
│   │   └── repository/       # データアクセス実装
│   ├── domain/       # ドメイン層
│   │   ├── entity/   # エンティティ
│   │   ├── repository/       # リポジトリインターface
│   │   └── service/  # ドメインサービス
│   ├── usecase/      # ユースケース層
│   │   ├── interactor/       # ビジネスロジック
│   │   └── port/     # ポート定義
│   ├── infrastructure/       # インフラ層
│   │   ├── database/ # データベース設定
│   │   ├── storage/  # ファイルストレージ
│   │   └── auth/     # 認証関連
│   └── lib/          # ライブラリ
│       └── graph/    # GraphQL関連
│           ├── generated/    # 自動生成ファイル
│           ├── loader/       # DataLoader (N+1対策)
│           └── schema/       # GraphQLスキーマ定義
│               └── schema.graphqls
└── config/           # 設定ファイル
```

## 設計指針

### GraphQL特化の原則
1. **スキーマファースト**: GraphQLスキーマから型とリゾルバーを自動生成
2. **型安全性**: gqlgenによる厳密な型チェック
3. **N+1問題対策**: DataLoaderパターンの実装
4. **関心の分離**: GraphQL層とビジネスロジック層の明確な分離

### アーキテクチャ層
- **Adapter層**: GraphQLリゾルバー、リポジトリ実装
- **Domain層**: ビジネスルールとエンティティ
- **UseCase層**: アプリケーションロジック
- **Infrastructure層**: 外部システム連携

テスタブルで保守性の高い、GraphQL APIに最適化されたアーキテクチャです。