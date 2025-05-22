package unusederror

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "unusederror",
	Doc:  "detects function calls that return error but whose result is ignored",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	errorType := types.Universe.Lookup("error").Type()
	fmt.Printf("errorType: %T\n", errorType)

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			exprStmt, ok := n.(*ast.ExprStmt)
			if !ok {
				return true
			}
			fmt.Printf("exprStmt: %T\n", exprStmt)

			callExpr, ok := exprStmt.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 呼び出している関数を解析
			// ObjectOf に渡すのは次の関数識別子
			// SelectorExpr:メソッドやパッケージ関数(ex. fmt.Println)
			// Ident:ユーザ定義関数（ex. some()）
			var obj types.Object
			switch fun := callExpr.Fun.(type) {
			case *ast.Ident:
				obj = pass.TypesInfo.ObjectOf(fun)
			case *ast.SelectorExpr:
				obj = pass.TypesInfo.ObjectOf(fun.Sel)
			default:
				// その他のケース（IndexExpr, TypeAssertなど）は無視
				return true
			}

			if obj == nil {
				return true
			}

			// Signature は（非組み込みの）関数またはメソッド型を表します。
			// sig には関数のシグネチャ情報が代入される
			sig, ok := obj.Type().Underlying().(*types.Signature)
			if !ok {
				return true
			}
			fmt.Printf("sig: %T\n", sig)

			results := sig.Results()
			for i := 0; i < results.Len(); i++ {
				// 戻り値の型が error 型と同じかどうかをチェック
				if types.Identical(results.At(i).Type(), errorType) {
					pass.Reportf(callExpr.Lparen, "function returns error, but result is ignored")
					break
				}
			}
			return true
		})
	}

	return nil, nil
}
