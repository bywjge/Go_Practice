package main

import "fmt"

func main() {
	m := make(map[string]int)
	n := map[string]int{}
	fmt.Println("n:", n)

	n["one"] = 1
	n["two"] = 2

	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // map[one:1 two:2]
	fmt.Println("n:", n)
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
