### 環境構築

```shell:
$ air init # .air.toml作成
$ go mod y-u-y-a/template-go # go.mod作成
```

### 使用ライブラリ

- Golang 1.18
- Air
- ent

- PostgreSQL

### テーブル追加方法

1. スキーマ定義ファイルを作成する

```shell:
$ go run entgo.io/ent/cmd/ent init [User]
```

2. ./ent/shema/配下にモデル定義ファイルが生成するので記述する
3. メソッドなどのファイルが生成する

```shell:
$ go generate ./ent
```

4. マイグレーションファイルを生成する

```shell:
$ make gen-migration
```

5. マイグレーションを実行する

```shell:
$ make migrate-up
```
