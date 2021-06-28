package main

import (
	"github.com/pydio/cells-linter/zapslices"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(zapslices.Analyzer)
}
