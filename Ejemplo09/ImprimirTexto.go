package main

import "fmt"

func main() {
	/*a := 20
	b := "Victor"*/
	var a int = 38
	var b string = "Victor"

	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v%%\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%+d\n", a)
	fmt.Printf("%o\n", a)
	fmt.Printf("%x\n", a)
	fmt.Printf("%X\n", a)
	fmt.Printf("%q\n", a)

	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", b)
}
