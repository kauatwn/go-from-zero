package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// CelebrateBirthday é um método associado a Person (Pointer Receiver): equivalente aos métodos de instância em C# que operam no "this".
// O asterisco (*) antes de Person significa que estamos operando em um ponteiro (referência).
func (p *Person) CelebrateBirthday() {
	p.Age++
}

func main() {
	// 1. Criando um ponteiro para Person (equivalente a "new Person()" no C#)
	// O operador & retorna o endereço de memória da struct.
	first := &Person{Name: "Alice", Age: 30}

	// 2. Comportamento de referência (Idêntico ao C#):
	// 'second' recebe o MESMO endereço de memória de 'first'. Nenhum dado é copiado.
	second := first
	second.Name = "Alice Updated"

	fmt.Println(first.Name) // Vai imprimir "Alice Updated", pois compartilham a referência!

	// 3. Chamando métodos com Pointer Receiver
	first.CelebrateBirthday()
	fmt.Println(first.Age) // Vai imprimir 31
}
