package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	fmt.Printf("GOTOOLCHAIN=%s\n", os.Getenv("GOTOOLCHAIN"))
	info, ok := debug.ReadBuildInfo() // 実行中のバイナリに埋め込まれたビルド情報を返す。
	if ok {
		fmt.Println("Go Version:", info.GoVersion) // GoVersion はバイナリをビルドした Go ツールチェーンのバージョン
	}
}
