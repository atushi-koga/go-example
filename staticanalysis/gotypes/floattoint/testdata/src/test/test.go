package test

import "fmt"

var sampleInt = 42
var sampleFloat = 3.14

func badCase() {
	i := int(sampleFloat)      // want "float-to-integer conversion may lose precision"
	fmt.Printf("i:= %v \n", i) // → 3

	i64 := int64(sampleFloat)      // want "float-to-integer conversion may lose precision"
	fmt.Printf("i64 = %v \n", i64) // → 3

	ui32 := uint32(sampleFloat)      // want "float-to-integer conversion may lose precision"
	fmt.Printf("ui32 = %v \n", ui32) // → 3
}

func goodCase() {
	f64 := float64(sampleInt) // ✅ int → float: 精度は落ちにくいのでOK
	fmt.Printf("f64 = %v \n", f64)

	_ = []byte("hello") // ✅ string → []byte: よく使われるのでOK
}
