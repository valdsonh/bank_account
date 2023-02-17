package main

import (
	"bank_account/clientes"
	"bank_account/contas"
	"fmt"
)

func main() {

	cliente1 := clientes.Titular{Nome: "Valdson", CPF: "123.456.789-00", Profissao: "Médico"}
	conta1 := contas.ContaCorrente{Titular: cliente1, NumeroAgencia: 123, NumeroConta: 123456}
	contaP1 := contas.ContaPoupanca{}

	// conta2 := contas.ContaCorrente{Titular: "Henrique", Saldo: 100}

	fmt.Println(conta1, contaP1)

	// status := conta2.Transferencia(-200, &conta1)

	// fmt.Println(status)
	// fmt.Println(conta1, conta2)

	// conta := ContaCorrente{}
	// conta.Titular = "Valdson"
	// conta.Saldo = 500

	// fmt.Println(conta.Saldo)

	// fmt.Println(conta.Saque(100))

	// status, valor := conta.Deposito(2000)

	// fmt.Println(status, valor)

}
