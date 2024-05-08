package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {

	for _, value := range list {
		fmt.Println(value)
	}
}

//parametro de  tipos y restricciones

type integer int

func Sum[T ~int | ~float64](nums ...T) T { //Puedo agragar el operrador |integer o usar la derivación de tipo de dato "~"
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

//Crear restricciones

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// constraints son restricciones que reutilizamos del paquete que descargamos
func Suma[T constraints.Integer | constraints.Float](nums ...T) T { //Puedo agragar el operrador |integer o usar la derivación de tipo de dato "~"
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

//Restricciones y operadores

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {

	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

// Estructuras genericas

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {

	PrintList("Nereo", "Neo", "Sinahi")
	PrintList(14, 30, 11)

	//parametro de  tipos y restricciones
	var nun1 integer = 100
	var nun2 integer = 300
	fmt.Println(Sum(11, 30, 14, 15))
	fmt.Println(Sum(nun1, nun2))

	//Crear restricciones
	fmt.Println(Suma[uint](11, 30, 14, 15))
	fmt.Println(Suma[float32](11.5, 30.5, 14.5, 15.5))

	//Restricciones y operadores
	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 1))
	fmt.Println(Includes(numbers, 11))

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	//Estructuras Genericas

	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"2", "Tenis", 100}

	fmt.Println(product1)
	fmt.Println(product2)
}
