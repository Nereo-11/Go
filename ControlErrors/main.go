package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero") //tambien se utiliza fmt.Errorf("")
	}
	return dividendo / divisor, nil
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}

	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)

}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
}

func main() {
	/*str := "123"
	num, err := strconv.Atoi(str)

	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Número:", num)
	*/

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(result)

	// Uso de Defer

	defer fmt.Println(3) // Si todos tienen defer se crearia una pila, el primero en entrar es el ultimo en salir
	fmt.Println(1)
	fmt.Println(2)

	file, err := os.Create("hola.texr") // Crea un archivo de texto
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Alex Roel")) // Defines lo que contendra el archivo de texto (txt)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Uso de panic y recover

	division(100, 10) // creamos función division y validateZero
	division(200, 25)
	division(100, 0)

	// Registro de errores

	log.SetPrefix("main():")
	log.Fatal("Soy un registro de errores")
	fmt.Println("Este es otro mensaje de registro")

	log.SetPrefix("main()")
	file, err = os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Soy un log")

}
