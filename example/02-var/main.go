package main

import (
	"fmt"
	"math"
)

func main() {

	// 声明变量的两种方式
	var a = "initial" // 默认推导类型

	var b, c int = 1, 2 // 也可以显式指定类型

	var d = true

	var e float64

	f := float32(e)
	//var f = float32(e)

	g := a + "foo"
	//var g = a + "foo"

	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant" // 常量
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
