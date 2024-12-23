# 概要
下記の値を設定してgoコマンドを実行すると、「toolchain not available」エラーになる。

- GOTOOLCHAIN:auto
- ローカルにインストールしたGoバージョン:1.21.0
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive:1.23
- go.mod の toolchain directive は無し

これは`GOTOOLCHAIN=auto`で go.mod の toolchain directive が無い場合、「goコマンドにバンドルされたバージョン」と「go directive のバージョン」を比較した結果、
「go directive のバージョン」が高いため。

# 動作確認
```
$ go version
go: downloading go1.23 (darwin/amd64)
go: download go1.23 for darwin/amd64: toolchain not available
```

```
$ go run main.go
go: downloading go1.23 (darwin/amd64)
go: download go1.23 for darwin/amd64: toolchain not available
```

```
$ go build main.go
go: downloading go1.23 (darwin/amd64)
go: download go1.23 for darwin/amd64: toolchain not available
```

```
$ docker build -t my-go-app .
[+] Building 4.1s (9/10)                                                                                           
 => [internal] load .dockerignore                                                                             0.0s
 => => transferring context: 2B                                                                               0.0s
 => [internal] load build definition from Dockerfile                                                          0.0s
 => => transferring dockerfile: 217B                                                                          0.0s
 => [internal] load metadata for docker.io/library/golang:1.21.0                                              2.8s
 => [auth] library/golang:pull token for registry-1.docker.io                                                 0.0s
 => [1/5] FROM docker.io/library/golang:1.21.0@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c947d  0.0s
 => [internal] load build context                                                                             0.0s
 => => transferring context: 2.13kB                                                                           0.0s
 => CACHED [2/5] WORKDIR /go/src                                                                              0.0s
 => [3/5] COPY . .                                                                                            0.0s
 => ERROR [4/5] RUN go mod download && go mod verify                                                          1.2s
------                                                                                                             
 > [4/5] RUN go mod download && go mod verify:
#0 0.150 go: downloading go1.23 (linux/amd64)
#0 1.179 go: download go1.23 for linux/amd64: toolchain not available
```
