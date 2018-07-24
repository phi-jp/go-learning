// 変数
package main

import "fmt"

func main() {
	// var msg string
	// msg = "hello world"
	// var msg = "hello world"
	msg := "hello world"
	fmt.Println(msg)

	a, b := 10, 15
	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
	fmt.Println(a, b, c, d)
}