package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

//Metodo para agregar tareas

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

//Metodo para marcar como completado

func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

//Metodo editar tareas

func (l *ListaTareas) editarTareas(index int, t Tarea) {
	l.tareas[index] = t
}

//Metodo Eliminar tareas

func (l *ListaTareas) eliminarTareas(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//Instancia

	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Println("Ingrese la opción")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese Nombre de Tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese Descripción de Tarea")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agrgada correctamente")

		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completada")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")

		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el indice de la tarea que desea editar")
			fmt.Scanln(&index)
			fmt.Println("Ingrese Nuevo Nombre de Tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese Nueva Descripción de Tarea")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTareas(index, t)
			fmt.Println("Tarea editada correctamente")

		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea editar")
			fmt.Scanln(&index)
			lista.eliminarTareas(index)
			fmt.Println("Tarea Eliminada")

		case 5:
			fmt.Println("Saliendo del Programa...")
			return
		default:
			fmt.Println("Opción invalida")

		}

		fmt.Println("Lista de Tareas:")
		fmt.Println("=====================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
	}

}
