//Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma função que receba
//um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).

package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
	Area   float64
}

func main() {
	var b, h float64

	fmt.Print("Base: ")
	fmt.Scan(&b)

	fmt.Print("Altura: ")
	fmt.Scan(&h)

	triangulo := Triangulo{
		Base:   b,
		Altura: h,
	}

	triangulo = CalculateTriangleArea(triangulo)
	fmt.Printf("Sua área é de %.2f", triangulo.Area)
}

func CalculateTriangleArea(triangulo Triangulo) Triangulo {
	triangulo.Area = triangulo.Base * triangulo.Altura / 2

	return triangulo
}
