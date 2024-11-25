package main

import (
	"unused"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unused.Analyzer) }
