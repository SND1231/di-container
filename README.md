# dip-test

使用方法

```
# docker-composeの作成、起動
docker-compose build
docker-compose up

# コードテスト実行
./test.sh

# 実際にS3に送られているか確認したい場合、AWSの設定をする
aws configure

# ファイルをS3に転送(同ディレクトリにあるsample.txtを試しに使用可能)
go run main.go {転送をしたいファイル} {バケット名} {オブジェクトキー(この名前でファイルが作成される)}

```
