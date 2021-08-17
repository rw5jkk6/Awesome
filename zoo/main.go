package main

import (
	"fmt"

	"./animals"
)

func main(){
	// 関数名が小文字で始まるとパッケージン変数にはならないのでエラーになる
	//fmt.Println(animals.elephantFeed())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())

	// 同じパッケージなので、関数が小文字で始まっても大丈夫
	// もちろん大文字でも大丈夫
	fmt.Println(appName())
}