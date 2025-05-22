package main

import (
	"errorimplement"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errorimplement.Analyzer) }
