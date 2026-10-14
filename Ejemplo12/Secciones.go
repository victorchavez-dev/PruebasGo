package main

import "fmt"

func main() {
	sPaisVecinos := []string{"Perú", "Ecuador", "Colombia", "Brasil", "Bolivia", "Chile"}
	sPaisLejanos := []string{"Argentina", "Uruguay", "Paraguay", "Venezuela"}

	fmt.Println(sPaisVecinos)
	fmt.Println("Largo", len(sPaisVecinos))
	fmt.Println("Cap", cap(sPaisVecinos))

	sPaisGuerra := make([]string, 4, 4)
	sPaisGuerra[0] = sPaisVecinos[5]
	sPaisGuerra[1] = sPaisVecinos[1]
	sPaisGuerra[2] = sPaisVecinos[2]
	sPaisGuerra[3] = sPaisVecinos[3]

	fmt.Println(sPaisGuerra)
	fmt.Println("Largo", len(sPaisGuerra))
	fmt.Println("Capacidad", cap(sPaisGuerra))

	fmt.Println("Agregar Slices")
	sPaisGuerra = append(sPaisGuerra, "Bolivia")

	fmt.Println(sPaisGuerra)
	fmt.Println("Largo", len(sPaisGuerra))
	fmt.Println("Capacidad", cap(sPaisGuerra))

	sPaisSur := append(sPaisVecinos, sPaisLejanos...)

	fmt.Println(sPaisSur)
	fmt.Println("Largo", len(sPaisSur))
	fmt.Println("Capacidad", cap(sPaisSur))
}
