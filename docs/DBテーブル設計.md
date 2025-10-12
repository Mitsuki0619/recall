# 🗄 DBテーブル設計

## 1. Userテーブル

| カラム名   | 型        | 制約             | 説明           |
| ---------- | --------- | ---------------- | -------------- |
| id         | UUID      | PK               | ユーザーID     |
| name       | TEXT      | NOT NULL         | ユーザー名     |
| email      | TEXT      | NOT NULL, UNIQUE | メールアドレス |
| avatar_url | TEXT      | NULL             | アイコンURL    |
| created_at | TIMESTAMP | DEFAULT now()    | 作成日時       |

## 2. UserProviderテーブル

| カラム名    | 型        | 制約           | 説明                                       |
| ----------- | --------- | -------------- | ------------------------------------------ |
| id          | UUID      | PK             | 連携ID                                     |
| user_id     | UUID      | FK -> User(id) | ユーザーID                                 |
| provider    | TEXT      | NOT NULL       | 連携プロバイダー名（Google/GitHub/LINE/X） |
| provider_id | TEXT      | NOT NULL       | プロバイダー側のユーザーID                 |
| created_at  | TIMESTAMP | DEFAULT now()  | 作成日時                                   |

## 3. Articleテーブル

| カラム名     | 型        | 制約             | 説明         |
| ------------ | --------- | ---------------- | ------------ |
| id           | UUID      | PK               | 記事ID       |
| user_id      | UUID      | FK -> User(id)   | 作成者ID     |
| url          | TEXT      | NOT NULL, UNIQUE | 記事URL      |
| title        | TEXT      | NOT NULL         | 記事タイトル |
| og_image_url | TEXT      | NULL             | OGP画像URL   |
| tags         | TEXT[]    | NULL             | タグIDリスト |
| created_at   | TIMESTAMP | DEFAULT now()    | 作成日時     |
| updated_at   | TIMESTAMP | DEFAULT now()    | 更新日時     |

## 4. Memoテーブル

| カラム名   | 型        | 制約              | 説明                 |
| ---------- | --------- | ----------------- | -------------------- |
| id         | UUID      | PK                | メモID               |
| article_id | UUID      | FK -> Article(id) | 対象記事ID           |
| title      | TEXT      | NOT NULL          | メモタイトル         |
| content    | TEXT      | NOT NULL          | メモ内容（Markdown） |
| created_at | TIMESTAMP | DEFAULT now()     | 作成日時             |
| updated_at | TIMESTAMP | DEFAULT now()     | 更新日時             |

## 5. Quizテーブル

| カラム名   | 型        | 制約              | 説明                                 |
| ---------- | --------- | ----------------- | ------------------------------------ |
| id         | UUID      | PK                | クイズID                             |
| article_id | UUID      | FK -> Article(id) | 対象記事ID                           |
| question   | TEXT      | NOT NULL          | 問題文                               |
| answer     | TEXT      | NOT NULL          | 正答または解答例                     |
| type       | TEXT      | NOT NULL          | 問題形式（true/false, choice, text） |
| created_at | TIMESTAMP | DEFAULT now()     | 作成日時                             |

## 6. Tagテーブル

| カラム名 | 型   | 制約     | 説明                     |
| -------- | ---- | -------- | ------------------------ |
| id       | UUID | PK       | タグID                   |
| name     | TEXT | NOT NULL | タグ名                   |
| color    | TEXT | NULL     | タグカラー（16進数など） |

## 7. ApiKeyテーブル

| カラム名   | 型        | 制約           | 説明           |
| ---------- | --------- | -------------- | -------------- |
| id         | UUID      | PK             | APIキーID      |
| user_id    | UUID      | FK -> User(id) | ユーザーID     |
| key        | TEXT      | NOT NULL       | OpenAI APIキー |
| created_at | TIMESTAMP | DEFAULT now()  | 作成日時       |
