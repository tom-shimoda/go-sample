◆ MariaDB使用方法 (参考: https://blog.s-style.co.jp/2019/06/4130/)
1. go mod init プロジェクト名
2. go get github.com/go-sql-driver/mysql
3. import以下を追加
```
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
```
