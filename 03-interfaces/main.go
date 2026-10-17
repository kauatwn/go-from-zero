package main

import "fmt"

// Notifier define o contrato (interface).
// Em Go, interfaces são implícitas: se o tipo possui o método, ele implementa a interface automaticamente.
type Notifier interface {
	Send(message string) string
}

type Email struct {
	Address string
}

// Send implementa a interface Notifier implicitamente (Pointer Receiver)
func (e *Email) Send(message string) string {
	return fmt.Sprintf("Sending e-mail from %s -> Message: %s\n", e.Address, message)
}

type SMS struct {
	ApiKey string
}

// Send também implementa a interface Notifier implicitamente
func (s *SMS) Send(message string) string {
	return fmt.Sprintf("Sending SMS using API Key %s -> Message: %s\n", s.ApiKey, message)
}

// NotifyUser é uma função polimórfica que aceita qualquer tipo que implemente a interface Notifier
func NotifyUser(n Notifier, message string) {
	fmt.Print(n.Send(message))
}

func main() {
	// Instanciando os serviços usando ponteiros (&)
	email := &Email{Address: "admin@system.com"}
	sms := &SMS{ApiKey: "12345-xyz"}

	message := "Hello, this is a notification!"

	// Polimorfismo sem precisar da palavra "implements"
	NotifyUser(email, message)
	NotifyUser(sms, message)
}
