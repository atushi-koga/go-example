package main

import (
	"called"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(called.Analyzer) }
