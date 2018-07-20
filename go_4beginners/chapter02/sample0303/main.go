// 関数
package main

import(
	"fmt"
)

func hello() {
	fmt.Println("hello")
}

func sum(i, j int) {
	fmt.Println(i + j)
}

func sum2(i, j int) int {
	return i + j
}

func main() {
	hello()
	sum(1, 2)
	fmt.Println(sum2(5, 10))
}
