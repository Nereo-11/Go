package main

import "fmt"

type Persona struct { //Declaración de estructura, funciona como un objeto
	nombre string
	edad   int
	correo string
}

func (persona *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", persona.nombre)

}

func main() {
	// Matrizes
	var matriz = [5]int{10, 20, 30, 40, 50}
	matriz[0] = 15

	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	for index, value := range matriz {
		fmt.Printf("Indice: %d, Valor: %d ", index, value)
	}

	var matrizBi = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matrizBi)

	// Rebanadas o Slices
	rebanada := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(rebanada)

	diaRebanada := rebanada[0:5]

	diaRebanada = append(diaRebanada, "Viernes", "Sabado")
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	diaRebanada = append(diaRebanada[0:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1) //Copiamos elementos de rebanada1 a rebanada2
	fmt.Println(rebanada2)

	//Mapas

	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors)         //Imprimir todos los elementos
	fmt.Println(colors["rojo"]) //Imprimir un solo elemento
	colors["negro"] = "#000000" //Agregar un nuevo elemento

	fmt.Println(colors)

	if valor, ok := colors["rojo"]; ok { //recuperar valor ("ok" es una verificación)
		fmt.Println("Si existe la clave")
		fmt.Println(valor)
	} else {
		fmt.Println("No existe esta clave")
	}

	delete(colors, "azul") //Eliminar elemento
	fmt.Println(colors)

	for clave, valor := range colors { // Iterar elementos
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}

	//Estructura, arriba de func main se agrega el tipo

	var p Persona //Creamos variable de tipo Persona
	p.nombre = "Alex"
	p.edad = 28
	p.correo = "alex@gmail.com"

	fmt.Println(p)

	p2 := Persona{"Diego", 29, "diego@gmail.com"} //Otra forma de declarar
	p2.edad = 30
	fmt.Println(p2)

	//Metodos y punteros
	var x int = 10
	var y *int = &x //puntero de x desde y

	fmt.Println(&x)
	fmt.Println(y)

	editar(&x)
	fmt.Println(x)

	//Metodo con mi estructura

	persona := Persona{"Jesus", 30, "jesus@gmail.com"}
	persona.diHola()
}

func editar(x *int) { //x recibe un puntero mediane "*int"

	*x = 20

}
