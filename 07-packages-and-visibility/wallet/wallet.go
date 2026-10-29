package wallet

// Wallet representa uma carteira.
// Como começa com W maiúsculo, ela é exportada (pública para outros pacotes).
type Wallet struct {
	balance float64 // balance começa com b minúsculo: é privado (visível apenas dentro do pacote wallet)
}

// NewWallet é uma função construtora exportada.
func NewWallet(initial float64) *Wallet {
	return &Wallet{balance: initial}
}

// Deposit altera o saldo de forma controlada. Exportada (D maiúsculo).
func (w *Wallet) Deposit(amount float64) {
	if amount > 0 {
		w.balance += amount
	}
}

// Balance expõe o saldo de forma segura, já que o campo 'balance' é privado. Exportada (B maiúsculo).
func (w *Wallet) Balance() float64 {
	return w.balance
}
