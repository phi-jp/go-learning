// 定数
package main

import(
	"fmt"
)

func main() {
	const hello string = "hello"
	// ちゃんとエラーになる
	// hello= "bye"
	fmt.Println(hello)
}
