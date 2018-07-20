// 関数は複数の値を返せる
package main

import(
	"fmt"
)

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	x, y := 3, 4
	x, y = swap(x, y);
	fmt.Println(x, y)

	// コンパイルエラーになる
	// x = swap(x, y)

	x, _ = swap(x, y)
	fmt.Println(x)
}
