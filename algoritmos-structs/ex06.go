//Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade". Escreva funções que permitam aumentar
//ou diminuir o salário do funcionário em uma determinada porcentagem e uma função que calcule o tempo de
//serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).

package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func TimeWorked(funcionario Funcionario) int {
	timeWorked := funcionario.Idade - 18

	return timeWorked
}

func IncreaseSalary(funcionario Funcionario) float64 {
	timeWorked := TimeWorked(funcionario)

	for i := 0; i < timeWorked; i++ {
		funcionario.Salario = funcionario.Salario + funcionario.Salario*5/100
	}

	return funcionario.Salario
}

func main() {
	funcionario1 := Funcionario{
		Nome:    "Tales",
		Salario: 2000.00,
		Idade:   19,
	}

	fmt.Print(IncreaseSalary(funcionario1))
}
