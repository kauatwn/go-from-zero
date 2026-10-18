package main

import "fmt"

func main() {
	// Slices (Equivalente aos Lists do C#)
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40)

	// O 'len' é o tamanho atual e o 'cap' é a capacidade alocada na memória.
	fmt.Println("Tamanho:", len(numbers), "Capacidade:", cap(numbers))
	fmt.Println()

	fmt.Println("Iterando sobre o slice:")
	for i, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	//// Range é mais idiomático em Go, mas também pode usar um loop tradicional.
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Printf("Index: %d, Value: %d\n", i, numbers[i])
	//}

	// Maps (Equivalente aos Dictionary<TKey, TValue> do C#)
	userAges := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// Adicionando ou atualizando valores no Map
	userAges["Charlie"] = 35
	fmt.Println()

	fmt.Println("Iterando sobre o map:")
	for name, age := range userAges {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// "Comma, ok" idiom do Go para buscar em Maps
	// Em C#, usa-se TryGetValue. Em Go, pega-se o valor e um booleano indicando se existe.
	age, exists := userAges["Alice"]
	if exists {
		fmt.Printf("\nAlice's age found: %d\n", age)
	}
}

// Em Go:
// - Slices ([]T) gerenciam arrays dinâmicos de forma eficiente por baixo dos panos.
// - Maps (map[K]V) são tabelas hash mutáveis passadas por referência implícita.
// - O 'range' substitui o foreach do C# para iterar sobre ambos.
