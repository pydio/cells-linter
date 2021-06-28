package zapslices

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "zapslices",
	Doc:  "Detect zap.Any calls with slices as arguments",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			be, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isPkgDot(be.Fun, "zap", "Any") {
				return true
			}

			if len(be.Args) < 2 {
				return true
			}

			if !isSlice(pass, be.Args[1]) {
				return true
			}

			pass.Reportf(be.Pos(), "zap.Any(slice) found %q",
				render(pass.Fset, be))
			return true

		})
	}

	return nil, nil
}

// helpers
// =======

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

func isPkgDot(expr ast.Expr, pkg, name string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	return ok && isIdent(sel.X, pkg) && isIdent(sel.Sel, name)
}

func isIdent(expr ast.Expr, ident string) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == ident
}

func isInteger(pass *analysis.Pass, expr ast.Expr) bool {
	t := pass.TypesInfo.TypeOf(expr)
	if t == nil {
		return false
	}

	bt, ok := t.Underlying().(*types.Basic)
	if !ok {
		return false
	}

	if (bt.Info() & types.IsInteger) == 0 {
		return false
	}

	return true
}

func isSlice(pass *analysis.Pass, expr ast.Expr) bool {
	t := pass.TypesInfo.TypeOf(expr)
	if t == nil {
		return false
	}

	_, ok := t.Underlying().(*types.Slice)
	return ok

}
