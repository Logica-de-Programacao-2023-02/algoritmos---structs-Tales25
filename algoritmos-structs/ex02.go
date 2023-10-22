//Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra
//struct com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa como
//parâmetro e imprima seu endereço completo.

package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

type Endereco struct {
	Estado   string
	Cidade   string
	Conjunto int
	Numero   int
}

func main() {
	pessoa := Pessoa{
		Nome:  "Tales",
		Idade: 19,
		Endereco: Endereco{
			Estado:   "DF",
			Cidade:   "Brasília",
			Conjunto: 7,
			Numero:   4,
		},
	}

	PrintAdress(pessoa)
}

func PrintAdress(pessoa Pessoa) {
	fmt.Println(pessoa.Endereco.Estado)
	fmt.Println(pessoa.Endereco.Cidade)
	fmt.Println(pessoa.Endereco.Conjunto)
	fmt.Println(pessoa.Endereco.Numero)
}
