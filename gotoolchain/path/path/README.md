# 概要
下記の値を設定してgoコマンドを実行すると、GOTOOLCHAINで指定したGoバージョンで実行される。

- GOTOOLCHAIN:go1.21.2+path
- ローカルにインストールしたGoバージョン:1.21.2
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive:1.16
- go.mod の toolchain directive:go1.21.0

これは`GOTOOLCHAIN=name+path`の場合、「nameに指定したバージョン」と「toolchain directive のバージョン」を比較した結果、
「nameに指定したバージョン」が高ければその値を採用するため。

# 動作確認
```
$ go version
go version go1.21.2 darwin/amd64
```

```
$ go run main.go
GOTOOLCHAIN=go1.21.2+path
Go Version: go1.21.2
```

```
$ go build main.go
$ ./main
GOTOOLCHAIN=go1.21.2+path
Go Version: go1.21.2
```

```
$ docker build -t my-go-app .
$ docker run --rm my-go-app
GOTOOLCHAIN=go1.21.2+path
Go Version: go1.21.2
```
