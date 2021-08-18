# zoo

- １つのディレクトリには１つのパッケージ名しか作ることができない
- 実行の仕方
  - go run *.go　または go run main.go animals.go 
  - go build
- テストの実行(-vは詳細の出力)
  - mainパッケージ　　　 go test -v
  - anmialsパッケージ　　　go test -v ./animals  
