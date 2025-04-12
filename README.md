# ToDoアプリ（Go + SQLite）

このアプリは、Go言語とSQLiteを使って作成されたシンプルなToDo管理アプリです。  
ユーザー登録、ログイン、タスクの追加・編集・削除といった基本機能を備えています。

## デモプレイ

![demo1](https://github.com/user-attachments/assets/0780266a-9212-4934-a0e5-c9a1e0cf1bfe)

---

![demo2](https://github.com/user-attachments/assets/b890a141-d681-4243-8d42-59da76085b60)

---

![demo3](https://github.com/user-attachments/assets/259570a8-fa92-40c0-b5ca-595dfa499241)

---

![demo4](https://github.com/user-attachments/assets/7684a9ec-f016-441d-86f5-ea1b2d26898d)

---

## 主な機能

- ユーザー登録・ログイン・ログアウト
- タスクの追加・編集・削除
- セッション管理によるログイン状態の保持
- ログイン前後で表示が変わる画面（public/private）

---

## 導入手順（ローカルで実行）

### 前提

- **Go（1.20以上）** がインストールされていること  
  インストールされていない場合：[Go公式サイト](https://golang.org/dl/) からインストールしてください

---

### 1. リポジトリをクローン

```bash
git clone https://github.com/Kazu-K0032/todo-app.git
cd todo-app
```

---

### 2. モジュールを初期化（依存パッケージの取得）

```bash
go mod tidy
```

---

### 3. `config/config.ini` の設定確認  

現在は以下の設定がされています。ポートなどご自身の環境に合わせて変更してください。

```ini
[web]
port = 8080
logfile = debug.log
static = app/views

[db]
driver = sqlite3
name = webapp.sql
```

---

### 4. アプリを実行

```bash
go run main.go
```

- 実行時、 `webapp.sql`・`debug.log`が生成されます。

---

### 5. ブラウザでアクセス

`config/config.ini` のデフォルトに合わせています。
ポート番号を変更した場合は、適宜変更してください。

```ini
http://localhost:8080
```

---

## ディレクトリ構成

```ini
todo-app/
├── app/
│   ├── controllers/     // ルーティングやハンドラー
│   ├── models/          // DBモデルやロジック
│   └── views/           // HTMLテンプレート
├── config/              // 設定ファイル（INIなど）
├── static/              // CSS/JS などの静的ファイル
├── go.mod
├── main.go
└── README.md
```

---

## 使用技術

- Go（net/http, html/template）
- SQLite（軽量なファイルベースDB）
- Bootstrap（簡単なスタイル適用）
- jQuery（簡単な動的処理）

---

## メモ・補足

- SQLiteのデータは `webapp.sql` に保存されます
- セッションはクッキーとDBで管理しています
- テンプレートは `layout.html` を共通としてパーツを切り替える構成です
