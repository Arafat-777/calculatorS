package shapes

type Rectangle struct {
	name          string
	Length, Width float64
}

func NewRectangle(length, width float64) Rectangle {
	return Rectangle{name: "Прямоугольник", Length: length, Width: width}
}

func (r Rectangle) GetName() string        { return r.name }
func (r Rectangle) CalculateArea() float64 { return r.Length * r.Width }
