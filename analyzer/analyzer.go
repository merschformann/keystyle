package analyzer

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "keystyle",
	Doc:  "Check the style of keys in maps.",
	Run:  run,
}
