package restrict

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// Analyzer は fmt.Println の使用を検出する静的解析ルールです
var Analyzer = &analysis.Analyzer{
	Name: "nofmtprintln",
	Doc:  "detects usage of fmt.Println and reports it",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// 関数呼び出しか？
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 選択式（fmt.Println のような形式）か？
			selector, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// パッケージ識別子を取得（fmt）
			pkgIdent, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if pkgIdent.Name == "fmt" && selector.Sel.Name == "Println" {
				pass.Reportf(callExpr.Pos(), "use of fmt.Println is not allowed")
			}

			return true
		})
	}
	return nil, nil
}
