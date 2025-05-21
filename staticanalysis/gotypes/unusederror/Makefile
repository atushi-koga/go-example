# ビルドしたバイナリを $GOPATH/bin/restrict として配置する
install:
	go install cmd/unusederror/unusederror.go
	ls -alt /Users/koga/go/bin/unusederror

# make install を実行した後、go vet -vettool=`which restrict` ./... で静的解析を実行できる
# 以下のコマンドは、動作確認可能なディレクトリである analysedcode 配下のコードに対して実行する例
analyze:install
	cd ../../analysedcode && go vet -vettool=`which unusederror` ./...

# 静的解析のテストコードを実行する
test:
	go test ./...
