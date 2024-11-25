package a

func f() {
	// チェックされるべき識別子名を使用しているため、この2行はDiagnosticが報告される。
	// 報告をアサートするにはテストデータの該当行にwantで始まるコメントを書く。
	// すると、want以降の""で括られている正規表現にマッチするDiagnosticが該当行で報告されているかチェックすることができる。
	var gopher int // want "identifier is gopher"
	print(gopher)  // want "identifier is gopher"

	var Gopher int // OK
	print(Gopher)  // OK
}
