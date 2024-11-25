package unused

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "unused",
	Doc:  doc,
	Run:  run, // 静的解析処理の本体
	Requires: []*analysis.Analyzer{ // 依存する Analyzer
		// 後のパスのためにASTトラバーサルを最適化するための Analyzer
		inspect.Analyzer,
		// golang.org/x/tools/go/analysis/passes パッケージの inspectパッケージは、
		// パッケージの抽象構文ツリーを探索するためのインスペクター (golang.org/x/tools/go/ast/inspector.Inspector) を提供するアナライザーを定義している

	},
}

const doc = "unused find unused identifyers"

func run(pass *analysis.Pass) (interface{}, error) {
	// pass.ResultOf には、依存する Analyzer の結果が格納されている。
	// inspect.Analyzer の結果は *inspector.Inspector 型でありこれを取得している。
	// *inspector.Inspector は、パッケージの抽象構文ツリーを検査(トラバース)するためのメソッドを提供する。
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}
	objects := map[types.Object][]*ast.Ident{}

	// inspect.Preorder:フィルターをかけて探索できる。
	// 第一引数の types が空でない場合、イベントの型ベースのフィルタリングが有効になる。
	// 第二引数の関数 f は、types スライスの要素と一致する型を持つノードに対してのみ呼び出される。
	// TODO: 関数定義・メソッド定義・パッケージ変数などに限定してPreorderでき用にしたものを別途作る
	// 関数・メソッド定義は (*ast.FuncDecl)(nil) 、 パッケージ変数や定数などの一般宣言は (*ast.GenDecl)(nil) でフィルタできそう
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			// 「_」を除外する理由
			// 「_」は空白識別子であり、どのスコープにも属さない無名の変数である。
			// 変数や値を明示的に「使わない」ことを示したり、関数の戻り値の一部や不要な値を無視するために使用される。
			// 使用例1: _, err := someFunction()
			// 使用例2:for _, v := range someSlice { // 何らかの処理 }
			// 空白識別子は本来「無視するため」に使用されるものであり、未使用かどうかの検出対象に含めるべきではないため、除外している。
			// そうでなければ上記の使用例1,2で「_」が未使用として検出されてしまう。
			if !ast.IsExported(n.Name) && n.Name != "_" {
				// 関数ブロック内で定義した変数名もここに入ってくる

				// pass.TypesInfo は types.Info 型で、型チェックの結果の型情報を持っている
				// types.Info は、型情報を保持するための構造体で、型情報は types.Object 型で保持される
				// types.Object は型情報を保持するためのインターフェース。パッケージ、定数、型、変数、関数 (メソッドを含む)、ラベルなどの名前付き言語エンティティを記述する。型情報は types.Var, types.Func, types.Const などの具体的な型で保持される。
				if o := pass.TypesInfo.ObjectOf(n); !skip(o) {
					objects[o] = append(objects[o], n)
				}
			}
		}
	})
	for o := range objects {
		if len(objects[o]) == 1 {
			n := objects[o][0]
			pass.Reportf(n.Pos(), "%s is unused", n.Name)
		}
	}

	return nil, nil
}

func skip(o types.Object) bool {
	if o == nil || o.Parent() == types.Universe {
		return true
	}
	switch o := o.(type) {
	case *types.PkgName:
		return true
	case *types.Var:
		if o.Pkg().Scope() != o.Parent() &&
			!(o.IsField() && !o.Anonymous() && isFieldInNamedStruct(o)) {
			return true
		}
	case *types.Func:
		// main
		if o.Name() == "main" && o.Pkg().Name() == "main" {
			return true
		}

		// init
		if o.Name() == "init" && o.Pkg().Scope() == o.Parent() {
			return true
		}

		// method
		sig, ok := o.Type().(*types.Signature)
		if ok {
			// インターフェースを実装している型のメソッドは未使用チェックから除外する
			// （インターフェース経由で使用し、実装を直接呼ばないケースがあるため）
			if recv := sig.Recv(); recv != nil {
				for _, i := range interfaces(o.Pkg()) {
					if i == recv.Type() ||
						(types.Implements(recv.Type(), i) && has(i, o)) {
						return true
					}
				}
			}
		}
	}

	return false
}

func interfaces(pkg *types.Package) []*types.Interface {
	var ifs []*types.Interface

	for _, n := range pkg.Scope().Names() {
		o := pkg.Scope().Lookup(n)
		if o != nil {
			i, ok := o.Type().Underlying().(*types.Interface)
			if ok {
				ifs = append(ifs, i)
			}
		}
	}

	return ifs
}

func has(intf *types.Interface, m *types.Func) bool {
	for i := 0; i < intf.NumMethods(); i++ {
		if intf.Method(i).Name() == m.Name() {
			return true
		}
	}
	return false
}

func isFieldInNamedStruct(v *types.Var) bool {
	structs := allNamedStructs(v.Pkg())
	for _, s := range structs {
		for i := 0; i < s.NumFields(); i++ {
			if s.Field(i) == v {
				return true
			}
		}
	}
	return false
}

func allNamedStructs(pkg *types.Package) []*types.Struct {
	var structs []*types.Struct

	for _, n := range pkg.Scope().Names() {
		o := pkg.Scope().Lookup(n)
		if o != nil {
			switch t := o.Type().(type) {
			case *types.Named:
				switch u := t.Underlying().(type) {
				case *types.Struct:
					structs = append(structs, u)
				}
			}
		}
	}

	return structs
}
