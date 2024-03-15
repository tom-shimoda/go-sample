MariaDB+phpmyadmin構築の参考元: https://keratoconu.hatenablog.com/entry/2023/02/20/173250

↑だけだとsessionフォルダへのアクセス権限不足によりphpエラーが出るためパーミッション設定を行う
https://konoyodehitori.hatenablog.com/entry/php-session_start-permission-denied/

phpmyadminログイン時、サーバーはdb-main (コンテナ名)にする


その他参考:
- TIMESTAMP列ではいったんUTCに変換されてから保存され、参照時にローカルタイムに再変換されるため、タイムゾーンにかかわらずUTCでデータが保存される
    https://qiita.com/cherubim1111/items/70cce9aaff500546f4d2

- JetBrains DataGripからwslのdocker上のDBに接続する方法
    Debianの場合コンソールから
    ```
    ip addr show
    ```
    を実行し、eth0のipアドレスをDataGripのHostに設定する
