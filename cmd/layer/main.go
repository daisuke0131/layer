package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"layer"
)

func main() { unitchecker.Main(layer.Analyzer) }
