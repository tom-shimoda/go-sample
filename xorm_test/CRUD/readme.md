参考: https://qiita.com/rin1208/items/fd536e95107aaf234bcf

0. dockerを起動しておく
```
sudo service docker start
```
1. docker/ にmysqlとテーブル表示用のphpmyddminのcomposeファイルがあるので
```
cd docker
docker compose up -d mysql
```
で起動

2. phpmyadmin起動
```
docker compose up -d phpmyadmin
```
ブラウザにてlocalhost:8000で接続

3. composeシャットダウン
```
docker compose down
```

Tips:
◆ composeファイルで構築したものをすべて削除するとき
(参考: https://qiita.com/suin/items/19d65e191b96a0079417)
```
docker compose down --rmi all --volumes --remove-orphans
```
