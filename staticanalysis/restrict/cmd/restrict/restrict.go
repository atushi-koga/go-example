package main

import (
	"restrict"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(restrict.Analyzer) }
