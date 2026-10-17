package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// CelebrateBirthday é um método associado a Person (Pointer Receiver).
// É equivalente aos métodos de instância em C# que operam no "this".
// O asterisco (*) antes de Person significa que estamos operando em um ponteiro (referência).
func (p *Person) CelebrateBirthday() {
	p.Age++
}

func main() {
	// Criando um ponteiro para Person (equivalente a "new Person()" no C#)
	// O operador & retorna o endereço de memória da struct.
	first := &Person{Name: "Alice", Age: 30}

	// Comportamento de referência (Idêntico ao C#):
	// 'second' recebe o MESMO endereço de memória de 'first'. Nenhum dado é copiado.
	second := first
	second.Name = "Alice Updated"

	// Vai imprimir "Alice Updated", pois compartilham a referência.
	fmt.Println(first.Name)

	first.CelebrateBirthday()
	fmt.Println(first.Age)
}
