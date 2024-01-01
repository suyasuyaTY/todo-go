# todo-go

golang+Docker(+AWS) のチュートリアル

MVC となるように実装する予定

利用ライブラリ

- [gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [golang-migration](https://github.com/golang-migrate/migrate)

# 今後の予定

todo-app の REST API

JWT 認証の実装

Redis?

# 環境構築

```
$ git clone git@github.com:suyasuyaTY/todo-go.git
$ cd ./todo-go
$ make up
$ make run
```

# migration

```
# マイグレーションファイル作成
$ migrate create -ext sql -dir db/migrations -seq (table-name)

# マイグレーション適用
$ make migrate-up

# リバースマイグレーション
$ make migrate-down
```
