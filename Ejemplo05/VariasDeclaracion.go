package main

import "fmt"

func main() {
	var a, b, c, d int = 1, 2, 3, 4

	fmt.Println(a, b, c, d)

	f, g, h, i := 1.2, 4.5, 6.7, 9.0
	fmt.Println(f, g, h, i)

	j, k := true, false
	fmt.Println(j, k)

	l, m := "Hello", "World"
	fmt.Println(l, m)

	n, o, p, q := 1, "Hello", true, 4.5
	fmt.Println(n, o, p, q)
}
