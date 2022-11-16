package main

import (
	"github.com/Deng-Xian-Sheng/gtimeFormatLinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
