package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Saque(valorDoSaque float64) string {
	saquePermitido := valorDoSaque <= c.saldo && valorDoSaque > 0

	if saquePermitido {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Deposito(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "O valor do depósito é menor que zero", c.saldo
	}

}

func (c *ContaCorrente) Transferencia(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Deposito(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {

	conta1 := ContaCorrente{titular: "Valdson", saldo: 300}
	conta2 := ContaCorrente{titular: "Henrique", saldo: 100}

	// fmt.Println(conta1, conta2)

	status := conta2.Transferencia(-200, &conta1)

	fmt.Println(status)
	fmt.Println(conta1, conta2)

	// conta := ContaCorrente{}
	// conta.titular = "Valdson"
	// conta.saldo = 500

	// fmt.Println(conta.saldo)

	// fmt.Println(conta.Saque(100))

	// status, valor := conta.Deposito(2000)

	// fmt.Println(status, valor)

}
