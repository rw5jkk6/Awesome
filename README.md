# zoo

- １つのディレクトリには１つのパッケージ名しか作ることができない
- 実行の仕方
  - go run *.go　または go run main.go animals.go 
  - go build
- テストについて
  - テストの仕様
    - ファイル名は xxx_test.goとする
    - 関数名には  func TestX(t *testing.T)とする
  - テストの実行(-vは詳細の出力)
    - mainパッケージ　　　 go test -v
    - anmialsパッケージ　　　go test -v ./animals  
