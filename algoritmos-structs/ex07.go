//Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som". Escreva funções que permitam
//modificar o som que o animal faz e uma função que imprima as informações do animal e o som que ele faz.

#PROFESSOR

// package main

// import "fmt"

// type Animal struct {
// 	Nome    string
// 	Especie string
// 	Idade   int
// 	Som     string
// }

// func main() {
// 	animal := Animal{
// 		Nome:    "Galinha",
// 		Especie: "Frango",
// 		Idade:   10,
// 		Som:     "Pó",
// 	}

// 	AnimalInfo(animal)

// 	animal = AlteraSom(animal, "Pó Pó Pó")

// 	AnimalInfo(animal)
// }

// func AlteraSom(a Animal, novoSom string) Animal {
// 	a.Som = novoSom
// 	return a
// }

// func AnimalInfo(a Animal) {
// 	fmt.Printf("O animal %s da espécie %s e com a idade de %d faz o som de %s\n",
// 		a.Nome,
// 		a.Especie,
// 		a.Idade,
// 		a.Som,
// 	)
// }

#EU

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	Nome    string
	Especie string
	Som     string
}

func main() {
	var newSound string
	scanner := bufio.NewScanner(os.Stdin)

	animal := Animal{
		Nome:    "Bob",
		Especie: "Cachorro",
		Som:     "Auau",
	}

	printInfo(animal)
	fmt.Print("Escolha um novo som: ")
	scanner.Scan()
	newSound = scanner.Text()

	//o animal do tipo Animal tem que ser alterado aqui, pq se não iria printar a mesma coisa
	animal = changeSound(animal, newSound)
	printInfo(animal)
}

func printInfo(animal Animal) {
	fmt.Printf("Um %s de nome %s faz '%s'\n",
		animal.Especie,
		animal.Nome,
		animal.Som,
	)
}

func changeSound(animal Animal, newSound string) Animal {
	animal.Som = newSound

	return animal
}
