package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"

	"rsc.io/quote"
)

// Declaración de constantes (Datos cuyos valores no cambian)
const Pi float32 = 3.14

const (
	x = 100
	y = 0b1010
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1 //Le asigna un valor a los siguientes dias
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola Mundo Neo11")
	fmt.Println(quote.Hello())

	var firstName, lastName string
	var age int

	/*
		Otras formas de declarar variables

		var (
			firstName, lastName string
			age                  int
		)



		firtsName, lastName, age := "Alex", "Roel", 11
	*/

	firstName = "Nereo"
	lastName = "Aranda"
	age = 11

	fmt.Println(firstName, lastName, age)

	//Covercion de tipo de datos

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))
	// Tambien...
	// fmt.Println(int32(integer16) + integer32)

	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i)

	n := 42
	s = strconv.Itoa(n)

	fmt.Println(s)

	name := "Neo"
	año := 2003

	fmt.Println("Ingrese su nombre")
	fmt.Scan(&name)
	fmt.Println("Ingrese su año")
	fmt.Scan(&año)

	mensaje := fmt.Sprintf("Hola me llamo %s y soy del año %d \n", name, año)
	fmt.Println(mensaje)

	// Operadores aritmeticas

	a := 10.0
	b := 11.0

	fmt.Println(a + b)

	lado := 0.0
	lado2 := 0.0
	area := 0.0
	hipotenusa := 0.0
	perimetro := 0.0

	fmt.Println("Ingrese la longitud de un lado del triangulo rectangulo")
	fmt.Scan(&lado)
	fmt.Println("Ingrese la longitud del segundo lado del triangulo rectangulo")
	fmt.Scan(&lado2)

	area = (lado * lado2) / 2
	hipotenusa = math.Sqrt(math.Pow(lado, 2) + math.Pow(lado2, 2))
	perimetro = lado + lado2 + hipotenusa

	fmt.Printf("\n Hipotenusa = %.2f \n", hipotenusa)
	fmt.Printf("\n Area = %.2f \n", area)
	fmt.Printf("\n Perimetro = %.2f \n", perimetro)

	// %.2f no permite elegir los decimales de los resultaos

	/*
			t := time.Now()
			hora := t.Hour()

			if hora < 12 {
				fmt.Println("Mañana")
			} else if hora < 17 {
				fmt.Println("Tarde")
			} else {
				fmt.Println("Noche")
		 	}
	*/
	os := runtime.GOOS

	switch os {
	case "Windows":
		fmt.Println("Windows")
	case "Linux":
		fmt.Println("Linux")
	case "Darwin":
		fmt.Println("macOS")
	default:
		fmt.Println("Otro OS")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana")
	case t.Hour() < 17:
		fmt.Println("Tarde")
	default:
		fmt.Println("Noche")

	}

	var h int
	for h <= 10 {
		fmt.Println(h)
		h++
	}

	for x := 1; x <= 10; x++ {
		fmt.Println(x)
		if x == 5 {
			break
		}
	}

	/*
		for x := 1; x <= 10; x++{

			if x == 5{
				continue
			}
			fmt.Println(x)
		}
	*/
	hello("Neo")
	saludo := msj("Nereo")
	fmt.Println(saludo)

	sum, res, mul := calculadora(4, 5)
	fmt.Println("La suma es", sum)
	fmt.Println("La resta es", res)
	fmt.Println("La suma es", mul)
}

func hello(name string) {
	fmt.Println("Hola ", name)
}

func msj(name string) string {
	return "Hola " + name
}

func calculadora(a, b int) (int, int, int) {
	suma := a + b
	resta := a - b
	multiplicacion := a * b
	return suma, resta, multiplicacion
}
