package main

import "fmt"

//sumar función de sumar dos numeros
func sumar(a, b int) int {
	return a + b
}

func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b
}

//Función principal
func main() {

	//Variables
	var a, b, c, r, suma int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función sumar
	suma = sumar(a, b)

	//Salida de datos
	fmt.Println("La suma es:", suma)

	//Llamar la función cociente
	c = cociente(a, b)

	//LLamar la función residuo
	r = residuo(a, b)

	//Salida de datos
	fmt.Println("Cociente:", c)
	fmt.Println("Residuo:", r)
}
