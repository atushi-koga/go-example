# 概要
Go1.21から導入された Go Toolchains の動きを確認するためのディレクトリ。

# 実行バージョンの選択
default とは、バンドルされたGoバージョンを指す。

```mermaid

graph TD
    classDef lightYellow fill:#FFFFE0,stroke:#000,stroke-width:2px,color:#000;
    A[GOTOOLCHAIN設定]
    A -->|local| B[バンドルされたツールチェーンを選択]
    A -->|name| C[その特定のツールチェーン。<br>PATHを見て無ければDL]
    A -->|path or name+path| D[異なるツールチェーンを許可<br>必要に応じて新しいGoバージョンを選択する。<br>**DLはせずに停止する**]
    A -->|auto or name+auto| E[異なるツールチェーンを許可<br>必要に応じて新しいGoバージョンを選択する。<br>**DLする**]

    F[go.modのtoolchain行を参照]
    D --> F
    E --> F
    F -->|toolchain tname| H[tname > default]
    F -->|toolchain行が無い| I[default >= go version]
    H -->|yes| J[tnameを選択]
    H -->|no| K[defaultを選択]
    I -->|yes| L[defaultを使用する]
    I -->|no| M[error:toolchain not available]
        
    Z[実行]
    B --> Z
    C --> Z
    J --> Z
    K --> Z
    L --> Z
    M --> Z

    class A,B,E,F,H,J,K,M,Z lightYellow;
```


# 参考
公式ドキュメント：[Go Toolchains](https://go.dev/doc/toolchain)
