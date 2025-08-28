package shapes

import "math"

type Circle struct {
	name   string
	Radius float64
}

func NewCircle(r float64) Circle {
	return Circle{name: "Круг", Radius: r}
}

func (c Circle) GetName() string        { return c.name }
func (c Circle) CalculateArea() float64 { return math.Pi * c.Radius * c.Radius }
