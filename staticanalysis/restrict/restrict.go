package restrict

import (
	"fmt"
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
	fmt.Println("----------start analysis----------")
	// pass.Files に解析対象のソースコードファイルが格納されている
	for _, file := range pass.Files {
		// ASTノードの構造を出力する
		//fmt.Print("--- AST Structure ---")
		// 出力形式は Go AST Viewer(https://yuroyoro.github.io/goast-viewer/) と同じ
		//err := ast.Print(pass.Fset, file)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "Error printing AST: %v\n", err)
		//}
		//fmt.Println("---------------------")

		// 深さ優先探索でファイル内の各ノードにアクセスしてこのクロージャを実行している
		ast.Inspect(file, func(n ast.Node) bool {
			// ast.Node に入るノードを調べる
			//fmt.Print("--- ast.Node start --- \n")
			//err2 := ast.Print(pass.Fset, n)
			//if err2 != nil {
			//	fmt.Fprintf(os.Stderr, "Error printing AST: %v\n", err2)
			//}
			//fmt.Print("--- ast.Node end --- \n")

			// 関数呼び出しか？
			// CallExpr:関数呼び出しを表す式ノード
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 選択式（fmt.Println のような形式）か？
			// *ast.SelectorExpr:フィールドやメソッドを参照する式
			selector, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// パッケージ識別子を取得（fmt）
			// *ast.Ident:識別子（変数名、関数名、型名など）を表すノード
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
