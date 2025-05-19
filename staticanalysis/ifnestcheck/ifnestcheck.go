package ifnestcheck

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// Analyzer は fmt.Println の使用を検出する静的解析ルールです
var Analyzer = &analysis.Analyzer{
	Name: "ifnestcheck",
	Doc:  "if文が3つ以上ネストになっている箇所をチェックする",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		v := &nestCheckVisitor{
			pass:      pass,
			nodeStack: []ast.Node{},
		}
		ast.Walk(v, file)
	}
	return nil, nil
}

// nestCheckVisitor は if文のネスト深さを記録し、深すぎるものを警告
type nestCheckVisitor struct {
	pass      *analysis.Pass
	depth     int
	nodeStack []ast.Node
}

func (v *nestCheckVisitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("Node type: %T, current depth: %d, stack size: %d\n", n, v.depth, len(v.nodeStack))

	if n != nil {
		v.nodeStack = append(v.nodeStack, n)
		if stmt, ok := n.(*ast.IfStmt); ok {
			v.depth++
			fmt.Printf("increment! depth = %v : %d\n", v.depth)

			if v.depth >= 3 {
				v.pass.Reportf(stmt.Pos(), "⚠️ if 文のネストが深すぎます（深さ: %d）", v.depth)
			}
		}
	} else {
		if len(v.nodeStack) > 0 {
			poppedNode := v.nodeStack[len(v.nodeStack)-1]
			v.nodeStack = v.nodeStack[:len(v.nodeStack)-1]

			if _, ok := poppedNode.(*ast.IfStmt); ok {
				v.depth--
				fmt.Printf("decrement! depth = %v : %d\n", v.depth)
			}
		} else {
			fmt.Println("Warning: Visit(nil) called with empty nodeStack")
		}
	}
	return v
}
