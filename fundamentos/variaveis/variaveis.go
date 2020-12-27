package main

import (
	"fmt"
	m "math" // Alias para o pacote math
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * m.Pow(raio, 2)

	fmt.Printf("O valor da area da circunferencia eh: %.2f\n", area)

	// Outras formas de declarar variaveis/constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a,b,c,d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, 'a', "hipopotamo" // Aspas simples equivale ao int na ASCII
	fmt.Println(g, h, i)
}