package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos

	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"
	fmt.Println(dias)

	//nuevo mapa

	estudiantes := make(map[string][]int)

	estudiantes["Victor"] = []int{13, 15, 16}
	estudiantes["Jose"] = []int{14, 13, 17}

	fmt.Println(estudiantes)

	//acceder a un valor

	fmt.Println(estudiantes["Victor"][1])

}
