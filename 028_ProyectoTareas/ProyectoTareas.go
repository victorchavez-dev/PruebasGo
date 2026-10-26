package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	Nombre      string
	Descripcion string
	Completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Metodo para marcar como completado
func (l *ListaTareas) mostrarCompletado(index int) {
	l.tareas[index].Completada = true
}

// Metodo editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Metodo eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar Tarea\n",
			"2. Marca Tarea como completado:\n",
			"3. Editar Tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir")

		fmt.Print("Ingrese Opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la Tarea:")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la Tarea:")
			t.Descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada exitosamente")
		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completado:")
			fmt.Scan(&index)

			lista.mostrarCompletado(index)
			fmt.Println("Tarea marcada como completada exitosamente")
		case 3:
			var index int
			var t Tarea

			fmt.Println("Ingrese el indice de la tarea a actualizar:")
			fmt.Scan(&index)
			fmt.Println("Ingrese el nuevo nombre de la tarea:")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la nueva descripción de la tarea:")
			t.Descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada exitosamente")
		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea a eliminar:")
			fmt.Scan(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada exitosamente")
		case 5:
			return
		default:
			fmt.Println("Opcion no valida")
		}

		fmt.Println("Lista de Tareas:")
		fmt.Println("=====================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s [%t]\n", i, t.Nombre, t.Descripcion, t.Completada)
		}
		fmt.Println("=====================================")
	}

}
