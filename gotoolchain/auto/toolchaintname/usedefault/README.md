# 概要
下記の値を設定してgoコマンドを実行すると、ローカルにインストールしたGoバージョンで実行される。

- GOTOOLCHAIN:auto
- ローカルにインストールしたGoバージョン:1.23.0
  - `which go`で「1.23.0」であることを確認する
- go.mod の go directive:1.16
- go.mod の toolchain directive:go1.22.0

これは`GOTOOLCHAIN=auto`の場合、「goコマンドにバンドルされたバージョン」と「toolchain directive のバージョン」を比較した結果、
高い方のバージョンである「goコマンドにバンドルされたバージョン」が採用されるため。

# 動作確認
```
$ go version
go version go1.23.0 darwin/amd64
```

```
$ go run main.go
GOTOOLCHAIN=auto
Go Version: go1.23.0
```

```
$ go build main.go
$ ./main
GOTOOLCHAIN=auto
Go Version: go1.23.0
```

```
$ docker build -t my-go-app .
$ docker run --rm my-go-app
GOTOOLCHAIN=auto
Go Version: go1.23.0
```
