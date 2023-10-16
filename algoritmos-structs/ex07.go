//Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som". Escreva funções que permitam
//modificar o som que o animal faz e uma função que imprima as informações do animal e o som que ele faz.

package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func main() {
	animal := Animal{
		Nome:    "Galinha",
		Especie: "Frango",
		Idade:   10,
		Som:     "Pó",
	}

	AnimalInfo(animal)

	animal = AlteraSom(animal, "Pó Pó Pó")

	AnimalInfo(animal)
}

func AlteraSom(a Animal, novoSom string) Animal {
	a.Som = novoSom
	return a
}

func AnimalInfo(a Animal) {
	fmt.Printf("O animal %s da espécie %s e com a idade de %d faz o som de %s\n",
		a.Nome,
		a.Especie,
		a.Idade,
		a.Som,
	)
}
