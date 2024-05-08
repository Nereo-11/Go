package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p Perro) Sonido() {
	fmt.Println(p.Nombre + "Hace Gua Gua")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + "Hace Miau Miau")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
