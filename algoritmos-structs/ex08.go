//Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço". Escreva uma função
//que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.

package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func main() {
	var sliceViagens []Viagem
	var viagem1, viagem2, viagem3, viagem4, mostExpensiveTrip Viagem

	viagem1 = Viagem{
		Origem:  "Brasília",
		Destino: "São Paulo",
		Data:    "26/10/2023",
		Preco:   3500,
	}

	viagem2 = Viagem{
		Origem:  "Brasília",
		Destino: "Juiz de Fora",
		Data:    "26/10/2023",
		Preco:   1750,
	}

	viagem3 = Viagem{
		Origem:  "Brasília",
		Destino: "Gramado",
		Data:    "26/10/2023",
		Preco:   5000,
	}

	viagem4 = Viagem{
		Origem:  "Brasília",
		Destino: "Orlando",
		Data:    "26/10/2023",
		Preco:   12500,
	}

	sliceViagens = []Viagem{viagem1, viagem2, viagem3, viagem4}

	mostExpensiveTrip = calculateExpensiveTrip(sliceViagens)
	fmt.Printf("A viagem mais cara é de %s - %s, com data %s e preço R$%.2f",
		mostExpensiveTrip.Origem,
		mostExpensiveTrip.Destino,
		mostExpensiveTrip.Data,
		mostExpensiveTrip.Preco,
	)
}

func calculateExpensiveTrip(sliceViagens []Viagem) Viagem {
	var mostExpensiveTrip Viagem
	var mostExpensive float64

	for _, viagem := range sliceViagens {
		if viagem.Preco > mostExpensive {
			mostExpensiveTrip = viagem
		}
	}

	return mostExpensiveTrip
}
