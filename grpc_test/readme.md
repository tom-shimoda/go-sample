◆ protoのコンパイルを行うためには以下が必要となる
```
sudo apt install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

◆ また自動生成されたファイルのimportで参照エラーが出る場合は以下を行うためには行う
```
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

◆ curlコマンドのようにgRPCのリクエストを送れるツール
```
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

◆ gRPCurlの使い方
1. サーバー起動
```
go run cmd/server/main.go
```
2. サーバーに実装されているサービス一覧取得
```
grpcurl -plaintext localhost:8080 list
```
3. サービスのメソッド一覧取得
```
grpcurl -plaintext localhost:8080 list myapp.GreetingService
```
4. メソッド呼び出し
```
grpcurl -plaintext -d '{"name": "hsaki"}' localhost:8080 myapp.GreetingService.Hello
```



Tips: 

◆ 不要なものをgo getした場合、
```
go mod tidy
```
で除去でき、go installした場合は、~/go/bin/ 以下のバイナリを削除する
