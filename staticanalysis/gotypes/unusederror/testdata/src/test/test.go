package test

import (
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
	some2() // want "function returns error, but result is ignored"
}

func badCase2() {
	fmt.Fprintln(os.Stdout, "Hello!") // want "function returns error, but result is ignored"
}

func hoge() error {
	return some2() // OK
}

func some2() error {
	_, e := fmt.Fprintln(os.Stdout, "Hello, Golang") // OK
	return e
}
