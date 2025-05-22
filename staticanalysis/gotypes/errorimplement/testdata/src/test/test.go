package test

type MyError struct{} // want "type MyError implements the error interface"

func (e MyError) Error() string { return "error" }

type MyError2 struct{} // want "type MyError2 implements the error interface"

func (e *MyError2) Error() string { return "error" }
