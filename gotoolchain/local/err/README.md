# 概要
下記の値を設定してgo buildコマンドを実行すると、「go.mod requires go >= 1.22.0」エラーになる。

- GOTOOLCHAIN=local
- ローカルにインストールしたGoバージョン=1.21.0
  - `which go`で「1.21.0」であることを確認する
- go.mod の go directive=1.22.0
- go.mod の toolchain directive=go1.22.1

これは `GOTOOLCHAIN=local`の場合にローカルにインストールしたGoバージョンを実行しようとするが、go directive のバージョンの方が高いため。

# 動作確認
```
$ go version
go version go1.21.0 darwin/amd64
```

```
$ go run main.go
go: go.mod requires go >= 1.22.0 (running go 1.21.0; GOTOOLCHAIN=local)
```

```
$ go build main.go
go: go.mod requires go >= 1.22.0 (running go 1.21.0; GOTOOLCHAIN=local)
```

```
$ docker build -t my-go-app .
[+] Building 3.0s (9/10)                                                                                                 
 => [internal] load .dockerignore                                                                                   0.0s
 => => transferring context: 2B                                                                                     0.0s
 => [internal] load build definition from Dockerfile                                                                0.0s
 => => transferring dockerfile: 219B                                                                                0.0s
 => [internal] load metadata for docker.io/library/golang:1.21.0                                                    2.7s
 => [auth] library/golang:pull token for registry-1.docker.io                                                       0.0s
 => [1/5] FROM docker.io/library/golang:1.21.0@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c947de45069  0.0s
 => [internal] load build context                                                                                   0.0s
 => => transferring context: 1.73kB                                                                                 0.0s
 => CACHED [2/5] WORKDIR /go/src                                                                                    0.0s
 => [3/5] COPY .. .                                                                                                 0.0s
 => ERROR [4/5] RUN go mod download && go mod verify                                                                0.2s
------
 > [4/5] RUN go mod download && go mod verify:
#0 0.140 go: go.mod requires go >= 1.22.0 (running go 1.21.0; GOTOOLCHAIN=local)
```
