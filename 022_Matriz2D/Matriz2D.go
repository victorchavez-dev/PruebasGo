package main

import "fmt"

func main() {
	f := 4
	c := 4

	matriz := make([][]int, f)
	for i := range f {
		matriz[i] = make([]int, c)
		for j := range c {
			matriz[i][j] = i*c + j + 1
		}
	}

	fmt.Println("Matriz Resultante:")
	for i := range f {
		for j := range c {
			fmt.Printf("%3d ", matriz[i][j])
		}
		fmt.Println()
	}
}
