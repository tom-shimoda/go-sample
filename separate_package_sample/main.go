package main

import (
	"fmt"
	"separate_package_sample/PackageA"
)

func foo() {
	fmt.Println("foo")
}

func main() {
	foo()
	PackageA.Bar() // 外部ファイルの関数を使用する場合、頭文字は大文字
}
