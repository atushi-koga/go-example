package floattoint

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// Analyzer は float型から整数型への型変換を検出する静的解析です。
var Analyzer = &analysis.Analyzer{
	Name: "floatcast",
	Doc:  "detects any float-to-integer conversions that may lose precision",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 型変換かどうか判定（T(x) の形）
			if len(callExpr.Args) != 1 {
				return true
			}

			// キャスト先の型
			castType := pass.TypesInfo.TypeOf(callExpr.Fun)
			if castType == nil {
				return true
			}

			// キャスト元の型
			arg := callExpr.Args[0]
			argType := pass.TypesInfo.TypeOf(arg)
			if argType == nil {
				return true
			}

			// float型 → 整数型 を検出
			if isFloat(argType) && isInteger(castType) {
				pass.Reportf(callExpr.Lparen, "float-to-integer conversion may lose precision")
			}

			return true
		})
	}

	return nil, nil
}

// isFloat は float32 または float64 を検出
func isFloat(t types.Type) bool {
	b, ok := t.Underlying().(*types.Basic)
	return ok && (b.Kind() == types.Float32 || b.Kind() == types.Float64)
}

// isInteger はすべての整数型（符号付き・符号なし）を検出
func isInteger(t types.Type) bool {
	b, ok := t.Underlying().(*types.Basic)
	return ok && (b.Info()&types.IsInteger != 0)
}
