package main

import (
	"fmt"
	"go-from-zero/07-packages-and-visibility/wallet"
)

func main() {
	// Criando uma carteira usando a função pública do pacote wallet
	myWallet := wallet.NewWallet(100.0)

	// Depositando fundos
	myWallet.Deposit(50.0)

	// Lendo o saldo através do método público
	fmt.Printf("Current Balance: $%.2f\n", myWallet.Balance())

	// Tentando acessar o campo 'balance' diretamente resultaria em erro de compilação, pois é privado.
	// myWallet.balance = 200.0
}

// Em Go:
// - Não existem palavras-chave como 'public', 'private' ou 'internal'.
// - Letra MAIÚSCULA na primeira letra = Exportado (público para fora do pacote).
// - Letra MINÚSCULA na primeira letra = Não exportado (privado, restrito ao pacote de origem).
// - O encapsulamento é puramente estrutural baseado em convenção de nomenclatura.
