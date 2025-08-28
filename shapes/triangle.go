package shapes

type Triangle struct {
	name         string
	Base, Height float64
}

func NewTriangle(base, height float64) Triangle {
	return Triangle{name: "Треугольник", Base: base, Height: height}
}

func (t Triangle) GetName() string        { return t.name }
func (t Triangle) CalculateArea() float64 { return 0.5 * t.Base * t.Height }
