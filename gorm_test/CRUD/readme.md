参考: https://zenn.dev/a_ichi1/articles/4b113d4c46857a

- gormはxormと比べてRelationが張れるのでgormのほうがよいかも
    https://future-architect.github.io/articles/20190926/
    → DB自体に外部キー設定がされてるわけではないっぽい？
    https://qiita.com/sky0621/items/c4e98ef735e9f735bd1b
    そもそもormでmigrationせず、別途goのmigrationツールを探したほうがいいかも
    https://developers.cyberagent.co.jp/blog/archives/41187/
    https://qiita.com/kishimoto828/items/179072276799c740a3eb

- gormのModelについて
    https://qiita.com/gold-kou/items/45a95d61d253184b0f33

- JetBrains DataGripからwslのdocker上のDBに接続する方法
    Debianの場合コンソールから
    ```
    ip addr show
    ```
    を実行し、eth0のipアドレスをDataGripのHostに設定する
