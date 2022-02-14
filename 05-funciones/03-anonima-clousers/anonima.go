package main

import (
	"fmt"
	"strings"
)

//clousers

func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
	fmt.Println(repeat3("Mundo"))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Victor"))
	fmt.Println(repeat5("Jose"))

	//funcion anonima autoejecutable
	/*	func() {
			fmt.Println("Hola desde la funcion anonima")
		}()
	*/

	// funcion anonima con variable
	/*myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
	}

	fmt.Println(myfunc("Victor"))
	*/

}
