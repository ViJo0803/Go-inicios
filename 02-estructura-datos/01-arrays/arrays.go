package main

import "fmt"

func main() {
	//arrays se define la cantidad de datos y el tipo

	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)

	//array con valores

	nombres := [3]string{"Victor", "Jose", "Carlos"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
		"Rosado",
	}

	fmt.Println(colores, len(colores))

	//indices definidos

	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euros",
	}

	fmt.Println(monedas, len(monedas))

	//subarray desde el indice 0 hasta el indice 3
	subMoneda := monedas[0:3]
	fmt.Println(subMoneda)

}
