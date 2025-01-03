# 概要
Go1.21から導入された Go Toolchains の動きを確認するためのディレクトリ。

# 実行バージョンの選択
次のフローチャートは、go build を実行した時、ビルドされるバージョン選択の仕組みを図示したもの。
（goコマンドの中でも go version を実行する場合は、go.modのgo directiveの影響を受けないため go build に限定する）
default とは、バンドルされたGoバージョンを指す。

```mermaid
graph TD
classDef lightYellow fill:#FFFFE0,stroke:#000,stroke-width:2px,color:#000;
A[GOTOOLCHAIN設定]
A -->|local| B[指定バージョン=バンドルされたバージョン]
A -->|name| C[指定したバージョンが PATH になければ DL]
A -->|path or name+path| F
A -->|auto or name+auto| F

subgraph Group1 [指定したバージョンのみ実行]
    B
    C
    B --> N[指定したバージョン > go directive]
    C --> N
    N -->|yes| O[指定したバージョンを選択]
    N -->|no| P[go directive未満のためエラー]
end

%% サブグラフでDとEから派生するノードを囲む
subgraph Group2 [必要に応じてGoバージョンを選択]
    F[go.modのtoolchain行を参照]
    F -->|toolchain tname| H[tname > default]
    F -->|toolchain行が無い| I[default >= go version]
    H -->|yes| Q[tnameを選択<br>PATHにあるか]
    Q -->|no| S[GOTOOLCHAIN設定]
    S -->|auto or name+auto| V[tnameをDLする]
    S -->|path or name+path| R[エラー]
    H -->|no| K[defaultを選択]
    I -->|yes| L[defaultを使用する]
    I -->|no| M[error:toolchain not available]
end

Z[実行]
Q -->|yes| Z
K --> Z
L --> Z
O --> Z
V --> Z

class A,B,C,E,F,H,I,J,K,L,M,N,O,P lightYellow;
```


# 参考
公式ドキュメント：[Go Toolchains](https://go.dev/doc/toolchain)
