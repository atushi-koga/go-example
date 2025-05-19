package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"ifnestcheck"
)

func main() { unitchecker.Main(ifnestcheck.Analyzer) }
