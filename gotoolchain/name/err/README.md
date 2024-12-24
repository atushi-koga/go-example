# 概要
下記の値を設定してgo buildコマンドを実行すると、「go.mod requires go >= 1.23」エラーになる。

- GOTOOLCHAIN:go1.22.0
- ローカルにインストールしたGoバージョン:1.21.0
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive:1.23
- go.mod の toolchain directive:go1.23.0

これは `GOTOOLCHAIN=name`の場合に GOTOOLCHAIN で指定したバージョンで実行しようとするが、go directive のバージョンの方が高いため。

# 動作確認
```
$ go version
go version go1.22.0 darwin/amd64
```

```
$ go run main.go
go: go.mod requires go >= 1.23 (running go 1.22.0; GOTOOLCHAIN=go1.22.0)
```

```
$ go build main.go
go: go.mod requires go >= 1.23 (running go 1.22.0; GOTOOLCHAIN=go1.22.0)
```

```
$ docker build -t my-go-app .
[+] Building 5.7s (8/10)                                                                                                 
 => [internal] load build definition from Dockerfile                                                                0.0s
 => => transferring dockerfile: 222B                                                                                0.0s
 => [internal] load .dockerignore                                                                                   0.0s
 => => transferring context: 2B                                                                                     0.0s
 => [internal] load metadata for docker.io/library/golang:1.21.0                                                    2.5s
 => [auth] library/golang:pull token for registry-1.docker.io                                                       0.0s
 => [1/5] FROM docker.io/library/golang:1.21.0@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c947de45069  0.0s
 => [internal] load build context                                                                                   0.0s
 => => transferring context: 3.11kB                                                                                 0.0s
 => CACHED [2/5] WORKDIR /go/src                                                                                    0.0s
 => [3/5] COPY .. .                                                                                                 0.0s
 => [4/5] RUN go mod download && go mod verify                                                                      3.2s
 => => # go: downloading go1.22.0 (linux/amd64)
 > [4/5] RUN go mod download && go mod verify:                                                                           
#0 0.151 go: downloading go1.22.0 (linux/amd64)
#0 9.576 go: go.mod requires go >= 1.23 (running go 1.22.0; GOTOOLCHAIN=go1.22.0)
```
