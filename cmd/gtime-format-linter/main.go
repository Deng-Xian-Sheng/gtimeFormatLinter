package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"gtimeFormatLinter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
