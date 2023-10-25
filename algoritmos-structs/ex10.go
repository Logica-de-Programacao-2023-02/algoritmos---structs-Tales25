//Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações". O campo "avaliações" deve ser
//um slice de inteiros representando as notas que o filme recebeu dos espectadores. Escreva funções que permitam
//adicionar ou remover avaliações do filme, calcular a média das avaliações e imprimir as informações
//do filme e sua média de avaliações.

package main

import "fmt"

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []float64
	Media      float64
}

func main() {
	filme := Filme{
		Titulo:  "Matrix",
		Diretor: "Tawe",
		Ano:     2000,
	}

	printFilmeInfo(filme)

	//preciso atribuir a funcao ao filme para o valor atualizar
	filme = addReview(filme)
	filme = calculateMediaReview(filme)
	printFilmeInfo(filme)

	filme = removeReview(filme)
	filme = calculateMediaReview(filme)
	printFilmeInfo(filme)
}

func printFilmeInfo(filme Filme) {
	fmt.Printf("O filme %s - %d do diretor %s tem média de %.1f\n",
		filme.Titulo,
		filme.Ano,
		filme.Diretor,
		filme.Media,
	)
}

func addReview(filme Filme) Filme {
	var review float64

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite a %d avaliação: ", i+1)
		fmt.Scan(&review)

		filme.Avaliacoes = append(filme.Avaliacoes, review)
	}

	return filme
}

func calculateMediaReview(filme Filme) Filme {
	var totalReview float64

	for _, review := range filme.Avaliacoes {
		totalReview += review
	}

	filme.Media = totalReview / float64(len(filme.Avaliacoes))

	return filme
}

func removeReview(filme Filme) Filme {
	var position int

	fmt.Print("Digite a posição da avaliação que deseja remover: ")
	fmt.Scan(&position)

	filme.Avaliacoes = append(filme.Avaliacoes[:position-1], filme.Avaliacoes[position:]...)

	return filme
}
