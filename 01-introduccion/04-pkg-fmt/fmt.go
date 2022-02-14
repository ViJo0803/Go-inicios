package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Victor"
	edad := 35

	// %s string  %d int %v no se sabe que se imprime
	// \n salto de linea

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad es %d \n", nombre, edad)
	fmt.Println(mensaje)

	// saber que tipo de dato es

	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	print("otro nombre:", nombre)
}
