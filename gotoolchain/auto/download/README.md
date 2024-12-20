# 概要
下記の値を設定してgoコマンドを実行すると、toolchain directiveで指定したGoバージョンで実行される。

- GOTOOLCHAIN:auto
- ローカルにインストールしたGoバージョン:1.21.0
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive:1.16
- go.mod の toolchain directive:go1.22.0

これは`GOTOOLCHAIN=auto`の場合、「goコマンドにバンドルされたバージョン」と「toolchain directive のバージョン」を比較した結果、
高い方のバージョンである「toolchain directive のバージョン」が採用されるため。

# 動作確認
```
$ go version
go: downloading go1.22.0 (darwin/amd64)
go version go1.22.0 darwin/amd64
```

```
$ go run main.go
GOTOOLCHAIN=auto
Go Version: go1.22.0
```

```
$ go build main.go
$ ./main
GOTOOLCHAIN=auto
Go Version: go1.22.0
```

```
$ docker build -t my-go-app .
[+] Building 11.2s (8/10)                                                                                        
=> [1/5] FROM docker.io/library/golang:1.21.0@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c94  0.0s
=> CACHED [2/5] WORKDIR /go/src                                                                            0.0s
=> [3/5] COPY . .                                                                                          0.0s
=> [4/5] RUN go mod download && go mod verify                                                              7.5s
=> => # go: downloading go1.22.0 (linux/amd64)
=> [5/5] RUN go build main.go                                                                              7.4s
=> exporting to image                                                                                      3.1s
=> => exporting layers                                                                                     3.1s
=> => writing image sha256:5d2777f0cf6aba98498a15b4b278a839ba19921a8c6461c0e2fde80239e3e050                0.0s
=> => naming to docker.io/library/my-go-app

$ docker run --rm my-go-app
GOTOOLCHAIN=auto
Go Version: go1.22.0
```
