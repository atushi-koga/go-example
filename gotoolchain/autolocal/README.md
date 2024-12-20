- GOTOOLCHAIN:auto
- ローカルにインストールしたGoバージョン:1.23.0
  - `which go`で「1.23.0」であることを確認する
- go.mod の go directive:1.16
- go.mod の toolchain directive:go1.22.0

上記の値を設定してgoコマンドを実行すると、「GOTOOLCHAIN=auto」により「goコマンドにバンドルされたバージョン(1.23.0)」と「toolchain directive(1.22.0)」を比較し、
「goコマンドにバンドルされたバージョン」の方が大きいため、そのバージョンで実行する。

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
