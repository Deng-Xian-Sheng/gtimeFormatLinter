package analyzer

import (
	"flag"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

var Analyzer = &analysis.Analyzer{
	Name:  "gtimeFormatLinter",
	Doc:   "gtime.Time is the time type of the Go Frame framework. The formal parameters of its Format method are completely different from those of the Format method in the standard library.",
	Flags: flag.FlagSet{},
	Run:   run,
}

func callMultiplexing(call *ast.CallExpr, pass *analysis.Pass, formatSymbols []string) {
	if call.Args == nil || len(call.Args) == 0 {
		return
	}
	for _, arg := range call.Args {
		if arg == nil {
			continue
		}
		bas, ok := arg.(*ast.BasicLit)
		if !ok {
			continue
		}
		var existence bool
		for _, formatSymbol := range formatSymbols {
			if strings.Contains(bas.Value, formatSymbol) {
				existence = true
			}
		}
		if existence {
			pass.Reportf(bas.Pos(), "Incorrect formal parameters of the Format method of gtime.Time")
		}
	}
}
func run(pass *analysis.Pass) (interface{}, error) {
	var formatSymbols = []string{"2006", "01", "02", "15", "04", "05"}
	inspect := func(node ast.Node) bool {
		fun, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if fun.Body == nil || fun.Body.List == nil || len(fun.Body.List) == 0 {
			return true
		}
		for _, stmt := range fun.Body.List {
			if stmt == nil {
				continue
			}
			expr, ok := stmt.(*ast.ExprStmt)
			if !ok {
				continue
			}
			if expr.X == nil {
				continue
			}
			call, ok := expr.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			if call.Fun == nil {
				continue
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			if sel.Sel == nil || sel.Sel.Name != "Format" {
				continue
			}
			if sel.X == nil {
				continue
			}
			switch sel.X.(type) {
			case *ast.Ident:
				id := sel.X.(*ast.Ident)
				if id.Obj == nil {
					continue
				}
				if id.Obj.Decl == nil {
					continue
				}
				obj, ok := id.Obj.Decl.(*ast.ValueSpec)
				if !ok {
					continue
				}
				if obj.Type == nil {
					continue
				}
				sel2, ok := obj.Type.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				if sel2.X == nil {
					continue
				}
				id2, ok := sel2.X.(*ast.Ident)
				if !ok {
					continue
				}
				if id2.Name != "gtime" {
					continue
				}
				if sel2.Sel == nil {
					continue
				}
				if sel2.Sel.Name != "Time" {
					continue
				}
				callMultiplexing(call, pass, formatSymbols)
			case *ast.CallExpr:
				call2, ok := sel.X.(*ast.CallExpr)
				if !ok {
					continue
				}
				if call2.Fun == nil {
					continue
				}
				sel2, ok := call2.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				if sel2.X == nil {
					continue
				}
				id, ok := sel2.X.(*ast.Ident)
				if !ok {
					continue
				}
				if id.Name != "gtime" {
					continue
				}
				callMultiplexing(call, pass, formatSymbols)
			}
		}
		return true
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}
	return nil, nil
}
