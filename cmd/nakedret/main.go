package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/zalenskivolt/nakedret"
)

const (
	DefaultLines = 7
)

func main() {
	analyzer := nakedret.NakedReturnAnalyzer(DefaultLines)
	singlechecker.Main(analyzer)
}
