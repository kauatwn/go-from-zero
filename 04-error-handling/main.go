package main

import (
	"errors"
	"fmt"
)

type Email struct {
	Address string
}

// Send retorna uma string E um erro
func (e *Email) Send(message string) (string, error) {
	if e.Address == "" {
		return "", errors.New("e-mail address cannot be empty")
	}

	result := fmt.Sprintf("Sending e-mail from %s -> Message: %s\n", e.Address, message)
	return result, nil // nil significa "sem erro" (equivalente ao null do C# para referências)
}

func TestEmail(email *Email, message string) {
	output, err := email.Send(message)

	if err != nil {
		fmt.Printf("Error caught: %s\n", err)
		return
	}

	fmt.Println(output)
}

func main() {
	valid := &Email{Address: "admin@system.com"}
	fmt.Println("Testing a valid e-mail:")
	TestEmail(valid, "Hello!")

	invalid := &Email{Address: ""}
	fmt.Println("Testing an invalid e-mail:")
	TestEmail(invalid, "Hello!")
}

// Em Go, não existe try/catch.
// Erros são valores comuns retornados pelas funções e checados manualmente com 'if err != nil'.
