package main

import (
	"checkgopher"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkgopher.Analyzer) }

func f() {
	var gopher int
	print(gopher)
}
