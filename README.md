## テスト駆動開発
- レッド
  - 目的のテストを書く
  - 失敗する
- グリーン
  - 失敗したテストの新しいコードを書く
- リファクタリング
  - コードを綺麗にする
  - 重複の除去    

## テストのコツ
- テストしやすいように設計、データ構造になっていないか
## PythonとGoのテストの比較
 ||Go|Python|
 |:-:|:-:|:-:|
 |file名にtestをつける|最後|前|
 |import|関数の引数|classを継承|
 |メソッド、関数名にtestの頭|大文字|小文字|
 
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
