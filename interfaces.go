package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometrica) {
	fmt.Println("Medidas: ", g)
	fmt.Println("Área: ", g.area())
	fmt.Println("Perímetro: ", g.perimetro())
}
