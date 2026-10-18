package main

import "fmt"

// Calculate retorna a soma e a subtração de dois inteiros usando variáveis de retorno nomeadas na assinatura da função.
// Equivalente a tuplas em C#.
func Calculate(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return // "Bare return": retorna automaticamente as variáveis nomeadas
}

// SumAll recebe um número variável de inteiros e retorna a soma total.
// É uma função variádica, permitindo flexibilidade na quantidade de argumentos.
// Equivalente a usar 'params int[] numbers' em C#.
func SumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// ProcessNumber recebe um inteiro e uma função que processa esse inteiro.
// É uma função de ordem superior (Higher-Order Function).
// Equivalente a passar um delegate ou Func<int, int> em C#.
func ProcessNumber(n int, operation func(int) int) int {
	return operation(n)
}

func main() {
	// Usando o Retorno Nomeado
	sum, sub := Calculate(10, 5)
	fmt.Printf("Sum: %d, Sub: %d\n", sum, sub)

	// Usando a Função Variádica
	total := SumAll(1, 2, 3, 4, 5)
	fmt.Printf("Total sum: %d\n", total)

	// Usando Função como Valor (passando uma função anônima)
	double := func(x int) int {
		return x * 2
	}

	result := ProcessNumber(5, double)
	fmt.Printf("Processed result: %d\n", result)
}
