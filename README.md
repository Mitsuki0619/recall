# 🧠 Recall - エンジニア向けドキュメント学習支援アプリ

エンジニアが日々読む公式ドキュメントや技術記事（Zenn・Qiita・Mediumなど）を、リーディングリストとして保存・整理し、各記事にMarkdownでメモを残せるアプリ。さらに、保存した記事やメモ内容をもとにAIが自動で復習問題を生成し、学習内容の定着を支援する。

## ✨ 主な機能

- 📝 **記事管理**: URL入力による記事登録とOGP自動取得
- 🏷️ **タグ管理**: タグ・カテゴリによる記事の整理
- 📖 **Markdownメモ**: 記事ごとに複数のMarkdownメモを作成・編集
- 🤖 **AI復習問題**: メモ内容を元にOpenAI APIが自動で復習問題を生成
- 🔍 **検索・フィルタ**: タイトル・URL・タグ・メモ内容での検索
- 🔐 **認証**: Supabase AuthによるSSO対応（Google・GitHub・LINE・X）

## 🏗️ 技術構成

| 項目           | 技術・サービス                                       |
| -------------- | ---------------------------------------------------- |
| フロントエンド | React（Next.js）・Expo（React Native）               |
| API形式        | GraphQL                                              |
| バックエンド   | AWS Lambda（Go製 GraphQL API）                       |
| API管理        | AWS API Gateway                                      |
| 認証           | Supabase Auth（SSO対応：Google / GitHub / LINE / X） |
| データベース   | Supabase PostgreSQL                                  |
| ストレージ     | Amazon S3（OGPキャッシュ・画像アップロード等）       |
| AI連携         | OpenAI API（問題生成）                               |
| デプロイ       | AWS SAM / Serverless Framework                       |

## 📁 プロジェクト構成

```
recall/
├── apps/
│   ├── mobile/         # Expo React Native アプリ
│   └── web/            # Next.js Webアプリ
├── packages/           # 共通パッケージ
├── docs/               # 設計ドキュメント
│   ├── 要件定義.md
│   ├── DBテーブル設計.md
│   └── 画面設計/
└── README.md
```

## 🚀 開発開始

### 前提条件

- Node.js 18以上
- pnpm 9.0.0以上

### セットアップ

```bash
# 依存関係のインストール
pnpm install

# 全アプリの開発サーバー起動
pnpm dev

# 特定のアプリのみ起動
pnpm dev --filter=web
pnpm dev --filter=mobile
```

### ビルド

```bash
# 全アプリのビルド
pnpm build

# 特定のアプリのみビルド
pnpm build --filter=web
```

### その他のコマンド

```bash
# リント実行
pnpm lint

# リント自動修正
pnpm lint:fix

# フォーマット実行
pnpm fmt

# 型チェック
pnpm check-types
```

## 📚 ドキュメント

詳細な設計資料は `docs/` ディレクトリを参照してください：

- [要件定義](./docs/要件定義.md)
- [DBテーブル設計](./docs/DBテーブル設計.md)
- [画面設計](./docs/画面設計/)

## 🎯 想定ユーザー

- 学習中の学生エンジニア
- 若手エンジニア・キャリア初期層
- 日々ドキュメントや記事を読む実務エンジニア

## 💡 提供価値

- 技術記事を保存し、メモと共に体系的に学べる
- AIによる復習で理解度を測定・強化できる
- 自分だけの「技術知識ベース」を構築可能
