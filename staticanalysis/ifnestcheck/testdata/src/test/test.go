package test

import "fmt"

func badCase1() {
	str := "abc"
	fmt.Println(str) // want "use of fmt.Println is not allowed"
}

// 以下のコメントでエラーを抑制できるように改良したい
//lint:ignore restrict reason
