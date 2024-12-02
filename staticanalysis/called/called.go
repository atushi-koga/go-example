package called

import (
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

// 呼び出しを検出したい関数やメソッドを指定するフラグ。 この変数にフラグの値が格納される。
// 使用例
// package function: pkgname.Func
// method: (*pkgname.Type).Method
// go vet -vettool=`which called` -called.funcs="log.Fatal" main.go
var flagFuncs string

var Analyzer = &analysis.Analyzer{
	Name: "called",
	Doc:  Doc,
	Run:  run, // go vet コマンドで指定されたパッケージごとに実行される関数。
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func init() {
	// コマンドラインフラグを登録する
	Analyzer.Flags.StringVar(&flagFuncs, "funcs", "", "function or method names which are restricted calling")
}

const Doc = "called find callings specified by called.funcs flag"

func run(pass *analysis.Pass) (interface{}, error) {
	if flagFuncs == "" {
		// フラグが指定されていない場合は何もせず正常終了する
		return nil, nil
	}

	fs := restrictedFuncs(pass, flagFuncs)
	if len(fs) == 0 {
		// フラグの形式が不正な場合は何もせず正常終了する
		return nil, nil
	}

	// pass.Report は静的解析中に警告を記録するための関数
	// Report フィールドにカスタム関数を代入することで挙動を変更できる
	// analysisutil.ReportWithoutIgnore では、「//lint:ignore called」がコメントされている場合は警告を記録しないようにしている
	pass.Report = analysisutil.ReportWithoutIgnore(pass)
	// buildssa.Analyzer を用いて得られた SSA(Static Single Assignment) 形式のデータから、解析対象のソースコード内で定義された関数を取得
	// ここでいう関数とは、パッケージ内で定義した関数（ビルトイン関数やインポートされた関数は含まれない）、メソッド（構造体や型エイリアスに関連付けられたもの）、匿名関数・クロージャを指す
	// SSAを用いることで、関数呼び出しの解析が容易になる。
	// 具体例：
	// 間接的な呼び出しの検出: f := fmt.Println
	// 条件分岐・ループでの追跡: var logFunc func(string); if cond { logFunc = log.Fatal } else {logFunc = log.Println}; logFunc("Message")
	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, sf := range srcFuncs {
		// sf.Blocks は関数内のBasic Blockのリスト。
		// Basic Block とは、1つの入口と1つの出口を持つコードの連続した命令（Instruction）の集まり。
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				for _, f := range fs {
					// Instruction で f が呼び出されているかどうかを判定する
					if analysisutil.Called(instr, nil, f) {
						// メッセージをフォーマットしつつ、pass.Report 関数を呼び出して警告を記録する
						pass.Reportf(instr.Pos(), "%s must not be called", f.FullName())
						break
					}
				}
			}
		}
	}

	return nil, nil
}

func restrictedFuncs(pass *analysis.Pass, names string) []*types.Func {
	var fs []*types.Func
	for _, fn := range strings.Split(names, ",") {
		ss := strings.Split(strings.TrimSpace(fn), ".")

		// 長さ2未満は不正な形式のためスキップする
		if len(ss) < 2 {
			continue
		}
		// pkgname.Func 形式（パッケージ関数）で取り出す
		f, _ := analysisutil.ObjectOf(pass, ss[0], ss[1]).(*types.Func)
		if f != nil {
			fs = append(fs, f)
			continue
		}

		if len(ss) < 3 {
			continue
		}
		// (*pkgname.Type).Method 形式（メソッド）で取り出す
		pkgname := strings.TrimLeft(ss[0], "(")
		typename := strings.TrimRight(ss[1], ")")
		if pkgname != "" && pkgname[0] == '*' {
			pkgname = pkgname[1:]
			typename = "*" + typename
		}
		// パッケージの型を取り出す
		typ := analysisutil.TypeOf(pass, pkgname, typename)
		if typ == nil {
			continue
		}
		// メソッドを取り出す
		m := analysisutil.MethodOf(typ, ss[2])
		if m != nil {
			fs = append(fs, m)
		}
	}

	return fs
}
