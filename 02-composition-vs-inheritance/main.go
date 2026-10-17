package main

import "fmt"

// Animal funciona como "Classe Base" conceitual do C#
type Animal struct {
	Name string
}

// Eat é um método associado ao Animal (Pointer Receiver)
func (a *Animal) Eat() {
	fmt.Printf("%s is eating.\n", a.Name)
}

// Dog funciona como a "Classe Filha" que herda comportamentos (composição em Go)
type Dog struct {
	// Embedding: apenas o tipo, sem nome de campo.
	// Isso faz com que Dog "herde" os campos e métodos de Animal (Field/Method Promotion).
	Animal
	Breed string
}

func (d *Dog) Bark() {
	fmt.Printf("%s (the %s) is barking!\n", d.Name, d.Breed)
}

// Cat funciona como a "Classe Filha" que herda comportamentos (composição em Go)
type Cat struct {
	Animal
	IsLazy bool
}

func (c *Cat) Meow() {
	fmt.Printf("%s says Meow!\n", c.Name)
}

func main() {
	//dog := &Dog{
	//	Animal: Animal{Name: "Rex"},
	//	Breed:  "Golden Retriever",
	//}

	// Syntax sugar: podemos inicializar o campo Animal diretamente, sem precisar nomeá-lo.
	dog := &Dog{
		Name:  "Rex",
		Breed: "Golden Retriever",
	}

	fmt.Println(dog.Name) // Output: Rex

	dog.Eat()  // Output: Rex is eating.
	dog.Bark() // Output: Rex (the Golden Retriever) is barking!

	cat := &Cat{
		Name:   "Fluffy",
		IsLazy: true,
	}

	fmt.Println(cat.Name) // Output: Fluffy

	cat.Eat()  // Output: Fluffy is eating.
	cat.Meow() // Output: Fluffy says Meow!
}

// Em Go, herança por hierarquia é proibida.
// A filosofia do Go é clara: Prefira composição em vez de herança (Composition over inheritance).
