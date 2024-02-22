package PackageA

import (
	"fmt"
)

func Bar() {	// 外部ファイルで使用される場合、頭文字は大文字
	fmt.Println("bar")
}
