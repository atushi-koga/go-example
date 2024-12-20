# 概要
下記の値を設定してgoコマンドを実行すると、ローカルにインストールしたGoバージョンで実行される。

- GOTOOLCHAIN=local
- ローカルにインストールしたGoバージョン=1.21.0
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive=1.16
- go.mod の toolchain directive=go1.22.0

これは`GOTOOLCHAIN=local`の場合、goコマンドにバンドルされたバージョンが実行されるため。

# 動作確認
```
$ go version
go version go1.21.0 darwin/amd64
```

```
$ go run main.go
GOTOOLCHAIN=local
Go Version: go1.21.0
```

```
$ go build main.go
$ ./main
GOTOOLCHAIN=local
Go Version: go1.21.0
```

```
$ docker build -t my-go-app .
$ docker run --rm my-go-app
GOTOOLCHAIN=local
Go Version: go1.21.0
```
