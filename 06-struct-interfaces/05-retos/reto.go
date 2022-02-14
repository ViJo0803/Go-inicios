package main

import (
	"fmt"
	"math"
)

type Geometrica interface {
	area() float64
	perimetro() float64
}

type cuadrado struct {
	lado float64
}
type circulo struct {
	radio float64
}

func (cua *cuadrado) area() float64 {
	return cua.lado * cua.lado
}

func (cua *cuadrado) perimetro() float64 {
	return cua.lado * 4
}

func (cir *circulo) area() float64 {

	return math.Pi * (cir.radio * cir.radio)
}

func (cir *circulo) perimetro() float64 {
	return 2 * math.Pi * cir.radio
}

func medidas(g Geometrica) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())

}
func main() {
	circulo1 := circulo{radio: 4}
	medidas(&circulo1)

	cuadrado1 := cuadrado{lado: 3}
	medidas(&cuadrado1)

}
