package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("function divide(): %v", err)
		}
	}()
	validateZero(divisor)
	return dividendo / divisor, nil
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func main() {
	log.SetPrefix("main():")
	archivo, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	log.SetOutput(archivo)
	log.Print("Soy un log")

	result, _ := divide(10, 5)
	fmt.Println("Resultado: ", result)

	result, _ = divide(10, 0)
	fmt.Println("Resultado: ", result)

	result, _ = divide(30, 5)
	fmt.Println("Resultado: ", result)

	str := "123f"
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Hubo un error al convertir el string a int: %v", err)
	}
	fmt.Println(num)

}
