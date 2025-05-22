package main

import (
	"floattoint"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(floattoint.Analyzer) }
