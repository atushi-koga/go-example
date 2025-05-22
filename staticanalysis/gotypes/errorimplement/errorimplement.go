package errorimplement

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "errorimpl",
	Doc:  "detects types that implement the error interface",
	Run:  run,
}

// 「型定義だけ」を調べたいので、ast.Inspect を使わない以下の形をとっている
func run(pass *analysis.Pass) (interface{}, error) {
	// error インターフェースを types.Universe から取得

	errorType := types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				// ユーザ定義の型（Named型）だけを対象に
				obj := pass.TypesInfo.Defs[typeSpec.Name]
				if obj == nil {
					continue
				}

				named, ok := obj.Type().(*types.Named)
				if !ok {
					continue
				}

				// 型とポインタ型の両方で error 実装チェック
				if types.Implements(named, errorType) || types.Implements(types.NewPointer(named), errorType) {
					pass.Reportf(typeSpec.Pos(), "type %s implements the error interface", typeSpec.Name.Name)
				}
			}
		}
	}
	return nil, nil
}
