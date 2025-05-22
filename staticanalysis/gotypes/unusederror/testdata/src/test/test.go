package test

import (
	"errors"
	"fmt"
	"os"
)

func goodCase1() {
	err := some2()          // OK
	_, _ = fmt.Println(err) // OK
}

func goodCase2() {
	_ = hoge() // OK
}

func badCase1() {
	// ユーザ定義関数
	some2() // want "function returns error, but result is ignored"
}

func badCase2() {
	// ビルトイン関数
	fmt.Fprintln(os.Stdout, "Hello!") // want "function returns error, but result is ignored"
}

func badCase3() {
	x := Hoge{}
	x.Hoge() // want "function returns error, but result is ignored"
}

func hoge() error {
	return some2() // OK
}

func some2() error {
	_, e := fmt.Fprintln(os.Stdout, "Hello, Golang") // OK
	return e
}

type Hoge struct {
}

func (h Hoge) Hoge() error {
	return errors.New("hoge")
}
