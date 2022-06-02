## テストのコツ
- テストしやすいように設計、データ構造になっていないか
## Goのパッケージ
- １つのディレクトリには１つのパッケージ名しか作ることができない
- ディレクトリとパッケージ名は同じにする
- 実行の仕方
  - go run *.go　または go run main.go animals.go 
  - go build
- テストについて
  - テストの仕様
    - ファイル名は xxx_test.goとする
    - 関数名には  func TestXxx(t *testing.T)とする
  - テストの実行(-vは詳細の出力)
    - mainパッケージ　　　 go test -v
    - anmialsパッケージ　　　go test -v ./animals  
## zoo
- スターティングGo cf31~38より
## awesomeproject
- 覚える
