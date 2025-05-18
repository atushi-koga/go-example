package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fmt.Println("------------start------------")
	// 各種文を表現する例（指定された順序で設定）
	expressions := []string{
		// Stmt
		`1 +`,                                  // *ast.BadStmt:
		`const a = 1`,                          // *ast.DeclStmt:
		`;`,                                    // *ast.EmptyStmt:
		`label: fmt.Println("test")`,           // *ast.LabeledStmt:
		`fmt.Println("Hello")`,                 // *ast.ExprStmt:
		`ch <- 42`,                             // *ast.SendStmt:
		`x++`,                                  // *ast.IncDecStmt:
		`x, y := 1, 2`,                         // *ast.AssignStmt:
		`go fmt.Println("goroutine")`,          // *ast.GoStmt:
		`defer fmt.Println("deferred")`,        // *ast.DeferStmt:
		`return x`,                             // *ast.ReturnStmt:
		`break`,                                // *ast.BranchStmt:
		`{ fmt.Println("inside block") }`,      // *ast.BlockStmt:
		`if x > 0 { fmt.Println("positive") }`, // *ast.IfStmt:
		`switch x { case 1: fmt.Println("one") }`,          // *ast.SwitchStmt:
		`switch x.(type) { case int: fmt.Println("int") }`, // *ast.TypeSwitchStmt:
		`select { case <-ch: fmt.Println("received") }`,    // *ast.SelectStmt:
		`for i := 0; i < 10; i++ { fmt.Println(i) }`,       // *ast.ForStmt:
		`for k, v := range m { fmt.Println(k, v) }`,        // *ast.RangeStmt:

		// Expr
		`1 +`,            // *ast.BadExpr:
		`(x + y)`,        // *ast.ParenExpr:
		`v.M`,            // *ast.SelectorExpr:
		`arr[0]`,         // *ast.IndexExpr:
		`T[int, string]`, // *ast.IndexListExpr:
		`slice[1:3]`,     // *ast.SliceExpr:
		`x.(int)`,        // *ast.TypeAssertExpr:
		`f(1, 2, 3)`,     // *ast.CallExpr:
		`*ptr`,           // *ast.StarExpr:
		`-100`,           // *ast.UnaryExpr:
		`x + 1`,          // *ast.BinaryExpr:
		`key: "value"`,   // *ast.KeyValueExpr:
		`"literal"`,      // *ast.BasicLit:
		`x`,              // *ast.Ident:
		`...`,            // *ast.Ellipsis:
		// Type Literal
		`[3]int`,                  // *ast.ArrayType:
		`[]string`,                // スライス型
		`struct{X int; Y string}`, // *ast.StructType:
		`func(x int) string`,      // *ast.FuncType:
		`interface{Method() int}`, // *ast.InterfaceType:
		`map[string]int`,          // *ast.MapType:
		`chan int`,                // *ast.ChanType:
		// Literal
		`1`,                                // 数値リテラル
		`"foo"`,                            // 文字列リテラル
		`func(x int) int { return x + 1 }`, // 無名関数リテラル
		`[3]int{1, 2, 3}`,                  // 配列リテラル
		`[]string{"a", "b", "c"}`,          // スライスリテラル
		`struct{X int}{X: 42}`,             // 構造体リテラル
		`map[string]int{"key": 1}`,         // マップリテラル

		// Comment
		// Field
		// Case
		// File
		// Package
	}

	// 各文を解析してASTを出力
	for _, expr := range expressions {
		// パース用に式を一時的な関数に包む
		source := fmt.Sprintf("package main\nfunc dummy() { %s }", expr)
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "", source, parser.AllErrors)
		if err != nil {
			fmt.Printf("Parse error: %v, node: %#v \n", err, node)
			continue
		}
		var stmts []ast.Stmt
		ast.Inspect(node, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "dummy" {
				stmts = fn.Body.List
				return false
			}
			return true
		})
		// 文ごとに探索を実行
		fmt.Printf("Expression: %s\n", expr)
		for _, stmt := range stmts {
			fmt.Printf("stmts: %#v \n", stmt)
			ast.Inspect(stmt, func(n ast.Node) bool {
				return true
			})
		}
		fmt.Println()
	}
}
