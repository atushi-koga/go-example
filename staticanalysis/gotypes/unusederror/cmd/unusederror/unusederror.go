package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"unusederror"
)

func main() { unitchecker.Main(unusederror.Analyzer) }
