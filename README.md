# dip-test

使用方法

```
# docker-composeの作成、起動
docker-compose build
docker-compose up

# コンテナ名(NAMES)を取得
docker ps

# コンテナに入る
docker exec -it {コンテナ名} sh

# テスト実行
make test

# 実際にS3に送られているか確認したい場合、AWSの認証情報を設定する
aws configure

# ファイルをS3に転送(同ディレクトリにあるsample.txtを試しに使用可能)
go run main.go {転送をしたいファイル} {バケット名} {オブジェクトキー(この名前でファイルが作成される)}

```
