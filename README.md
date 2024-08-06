# mindfulbooks
# sample_api
## OpenAPI 命名規則
- 並び順はアルファベット順とする
- パスは、ケバブケース
- componentsは、スネークケース
- propertiesは、スネークケース
- デフォルト値がある場合は、デフォルト値を設定する
- リクエストボディは、components/requestBodiesに定義
- パスパラメータ、クエリパラメータは、components/parametersに定義
- レスポンスは、components/responsesに定義

## フォルダ構成

- controllers
  API 層から呼ばれる。
  サービス層の処理を呼び出す。
- models
  データ構造の定義
- repositories
  サービス層から呼ばれる。
  データベースの操作をする処理を書く。
- api
  main から利用する
  サービスとコントローラーの作成
  パスと関数の紐付け
- services
  コントローラー層から呼ばれる。
  リポジトリ層の処理を呼び出す。
- コントローラー層を介してユーザーの HTTP リクエスト・レスポンスとやりとりをする機能
  - ユーザーが送信した HTTP リクエストに含まれていたデータを受け取る
  - ユーザーに返す HTTP レスポンスに必要なデータを返す
- レポジトリ層を介してデータベースを扱う機能で
  - SQL クエリを含むレポジトリ層の関数を呼び出す
  - 呼び出したレポジトリ層の関数から、データベースから取得したデータを受け取る

  ## DB初期化
  ```
  docker-compose up -d
  ```

PGPASSWORDをつけると、一時的に環境変数を設定できる
```zsh
PGPASSWORD=password
```

  ```
PGPASSWORD=password psql -h  127.0.0.1 -U test -d mindfulbooks_db -f ./db/createTable.sql
```