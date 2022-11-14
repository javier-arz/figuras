package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (circulo *Circulo) area() float64 {
	return math.Pi * math.Pow(circulo.Radio, 2)
}

func (circulo *Circulo) perimetro() float64 {
	return math.Pi * circulo.Radio * 2
}
