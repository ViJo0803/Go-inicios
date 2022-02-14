package main

import "fmt"

//struct persona

type Persona struct {
	nombre string
	edad   int
}

//crear metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\n Edad: %d\n", p.nombre, p.edad)
}

//metodo para modificar
func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre

}

//Herencia
type Empleado struct {
	Persona
	sueldo float32
}

//crear objetos desde la estructura
func main() {
	p1 := Persona{"Victor", 35}
	p1.imprimir()
	//fmt.Println(p1)
	//p1.nombre = "jose"
	p1.editNombre("Jose")
	p1.imprimir()
	//fmt.Println(p1)
	p2 := Persona{nombre: "Juan", edad: 25}
	//fmt.Println(p2)
	p2.imprimir()
	p2.editNombre("Carlos")
	p2.imprimir()

	//herencia

	emp1 := Empleado{
		sueldo: 500,
	}

	emp1.nombre = "Pedro"
	emp1.edad = 30
	emp1.imprimir()
	fmt.Println(emp1)
}
