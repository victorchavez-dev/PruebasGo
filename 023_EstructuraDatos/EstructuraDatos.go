package main

import "fmt"

type Tarea struct {
	id     int
	nombre string
}

func main() {
	nTarea := 10

	listaTareas := make([]Tarea, nTarea)
	listaTareas[0].id = 1
	listaTareas[0].nombre = "Compras"
	listaTareas[1] = Tarea{
		id:     2,
		nombre: "Limpieza",
	}
	listaTareas[2] = Tarea{id: 3, nombre: "Correr"}
	listaTareas[3] = Tarea{id: 4, nombre: "Estudiar"}
	listaTareas[4] = Tarea{id: 5, nombre: "Programar"}
	listaTareas[5] = Tarea{id: 6, nombre: "Cenar"}
	listaTareas[6] = Tarea{id: 7, nombre: "Dormir"}
	listaTareas[7] = Tarea{id: 8, nombre: "Despertar"}
	listaTareas[8] = Tarea{id: 9, nombre: "Comer"}
	listaTareas[9] = Tarea{id: 10, nombre: "Beber"}

	fmt.Println("Lista de Tareas:")
	for _, tarea := range listaTareas {
		fmt.Printf("ID: %d, Nombre: %s\n", tarea.id, tarea.nombre)
	}

	fmt.Println("Memoria de estructura liberada")

}
