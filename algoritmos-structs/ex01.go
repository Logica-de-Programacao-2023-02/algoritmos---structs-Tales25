//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e calcule
//a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.

package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Raio float64
	Area float64
}

func CalculateArea(circle Circulo) Circulo {
	circle.Area = circle.Raio * circle.Raio * math.Pi

	return circle
}
func main() {
	var raio float64

	fmt.Print("Digite o raio de um Círculo: ")
	fmt.Scan(&raio)

	circle := Circulo{
		Raio: raio,
	}

	circle = CalculateArea(circle)
	fmt.Printf("Sua área é de %.2f", circle.Area)
}
