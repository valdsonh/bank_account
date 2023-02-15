package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	conta := ContaCorrente{titular: "Valdson", numeroAgencia: 589, numeroConta: 123456, saldo: 150.0}

	conta2 := ContaCorrente{"Bruna", 222, 111222, 225.75}

	fmt.Println(conta)
	fmt.Println(conta2)
}
