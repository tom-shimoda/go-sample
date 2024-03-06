1. Docker/mysql にmysqlのcomposeファイルがあるので
```
docker compose up -d mysql
```
で起動

2. mysqlコンテナ接続用コンテナがあるので
```
docker compose run cli
```
で必要に応じて接続

(参考: https://zenn.dev/mstn_/articles/ad5d7c7ad7e2d6)
