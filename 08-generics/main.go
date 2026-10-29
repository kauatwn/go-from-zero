package main

import (
	"cmp"
	"fmt"
)

// PrintElement é uma função genérica com 'any' (aceita qualquer tipo).
func PrintElement[T any](element T) {
	fmt.Printf("Element value: %v\n", element)
}

// Box é uma struct que pode armazenar qualquer tipo de item.
type Box[T any] struct {
	Item T
}

func (b *Box[T]) GetItem() T {
	return b.Item
}

// FindIndex é uma função genérica com 'comparable' que restringe a tipos que suportam comparação (==, !=).
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// Number é um Type Set que aceita int, float64 e qualquer tipo customizado baseado neles.
// O operador til (~) permite que tipos customizados como CustomInt (baseado em int) sejam aceitos.
type Number interface {
	~int | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// Max é uma função genérica que usa 'cmp.Ordered' para tipos que podem ser ordenados (int, float64, string, etc.).
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// CustomInt é um tipo customizado baseado em int.
// Criando um tipo customizado para testar o operador ~
type CustomInt int

func main() {
	PrintElement(42)         // Substitution: T → int
	PrintElement("Hello Go") // Substitution: T → string

	// Testando Struct Genérica
	intBox := &Box[int]{Item: 100}
	fmt.Printf("Int Box Item: %d\n", intBox.GetItem())

	// Testando comparable
	numbers := []int{10, 20, 30, 40}
	index := FindIndex(numbers, 30)
	fmt.Printf("Index found: %d\n", index)

	x := CustomInt(10)
	y := CustomInt(20)
	fmt.Printf("Sum of CustomInt: %v\n", Add(x, y))

	// Testando cmp.Ordered com números e strings
	fmt.Printf("Max int: %d\n", Max(15, 42))
	fmt.Printf("Max string: %s\n", Max("apple", "banana"))
}

// Em Go:
// - Usamos colchetes [T any] para definir tipos genéricos (semelhante ao <T> do C#).
// - 'any' aceita tudo, 'comparable' restringe a comparações de igualdade, e 'cmp.Ordered' a ordenação.
// - O operador til (~) permite que Type Sets aceitem tipos customizados criados sobre primitivos.
// - Conversões de tipo explícitas como CustomInt(20) são obrigatórias devido à tipagem estrita do Go.
