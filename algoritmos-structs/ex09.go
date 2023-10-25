//Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas". O campo "notas" deve ser um slice de float64,
//representando as notas que o aluno tirou em uma determinada disciplina. Escreva funções que permitam adicionar
//ou remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.

package main

import "fmt"

type Aluno struct {
	Nome  string
	Notas []float64
	Media float64
}

func main() {
	aluno := Aluno{
		Nome: "Tales",
	}

	printAlunoInfo(aluno)

	//atribuo as funções para aluno para que o aluno da main seja atualizado
	aluno = addNotas(aluno)
	aluno = calculateMedia(aluno)
	printAlunoInfo(aluno)

	aluno = removePiorNota(aluno)
	aluno = calculateMedia(aluno)
	printAlunoInfo(aluno)
}

func printAlunoInfo(aluno Aluno) {
	fmt.Printf("Com as notas %.2f aluno %s tem média %.2f\n", aluno.Notas, aluno.Nome, aluno.Media)
}

func addNotas(aluno Aluno) Aluno {
	var nota float64

	for i := 0; i < 4; i++ {
		fmt.Printf("Digite a %d nota: ", i+1)
		fmt.Scan(&nota)

		aluno.Notas = append(aluno.Notas, nota)
	}

	return aluno
}

func removePiorNota(aluno Aluno) Aluno {
	var piorNota float64
	piorNota = aluno.Notas[0]

	for i := 0; i < len(aluno.Notas); i++ {
		if aluno.Notas[i] < piorNota {
			piorNota = aluno.Notas[i]
		}
	}

	for j := 0; j < len(aluno.Notas); j++ {
		if piorNota == aluno.Notas[j] {
			aluno.Notas = append(aluno.Notas[:j], aluno.Notas[j+1:]...)
			break
		}
	}

	return aluno
}

func calculateMedia(aluno Aluno) Aluno {
	var notaTotal float64

	for _, nota := range aluno.Notas {
		notaTotal += nota
	}

	aluno.Media = notaTotal / float64(len(aluno.Notas))

	return aluno
}
