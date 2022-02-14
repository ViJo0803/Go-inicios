package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

const mensaje = "Hola desde mi constantes"

func funcionPrivada() {
	fmt.Println("Hola desde la funcion privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
