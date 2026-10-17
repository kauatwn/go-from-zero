package main

import "fmt"

// Animal funciona como "Classe Base" conceitual do C#
type Animal struct {
	Name string
}

// Eat é um método associado ao Animal (Pointer Receiver)
func (a *Animal) Eat() string {
	return fmt.Sprintf("%s is eating.", a.Name)
}

// Dog funciona como a "Classe Filha" que herda comportamentos (composição em Go)
type Dog struct {
	// Embedding: apenas o tipo, sem nome de campo.
	// Isso faz com que Dog "herde" os campos e métodos de Animal (Field/Method Promotion).
	Animal
}

func (d *Dog) Bark() string {
	return fmt.Sprintf("%s is barking!", d.Name)
}

// Cat funciona como a "Classe Filha" que herda comportamentos (composição em Go)
type Cat struct {
	Animal
}

func (c *Cat) Meow() string {
	return fmt.Sprintf("%s says meow!", c.Name)
}

func main() {
	//dog := &Dog{Animal: Animal{Name: "Rex"}}

	// Syntax sugar: podemos inicializar o campo Animal diretamente, sem precisar nomeá-lo.
	dog := &Dog{Name: "Rex"}

	fmt.Println(dog.Name)
	fmt.Println(dog.Eat())
	fmt.Println(dog.Bark())
	fmt.Println()

	cat := &Cat{Name: "Fluffy"}

	fmt.Println(cat.Name)
	fmt.Println(cat.Eat())
	fmt.Println(cat.Meow())
}

// Em Go, herança por hierarquia é proibida.
// A filosofia do Go é clara: Prefira composição em vez de herança (Composition over inheritance).
