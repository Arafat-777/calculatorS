package shapes

type Square struct {
	name string
	Side float64
}

func NewSquare(side float64) Square {
	return Square{name: "Квадрат", Side: side}
}

func (s Square) GetName() string        { return s.name }
func (s Square) CalculateArea() float64 { return s.Side * s.Side }
