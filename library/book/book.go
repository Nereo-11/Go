package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string // Un identificador privado comienza con minuscula, si comienza en mayuscula es publico o exportado
	author string
	pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %\n Author: %\nPages: %d\n", b.title, b.author, b.pages)
}

// Composición

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %\n Author: %\nPages: %d\n Editorial: %s\n Nivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}

//Polimorfismo  type Printable, Print
