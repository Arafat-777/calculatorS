package shapes

type Parallelogram struct {
	name         string
	Base, Height float64
}

func NewParallelogram(base, height float64) Parallelogram {
	return Parallelogram{name: "Параллелограмм", Base: base, Height: height}
}

func (p Parallelogram) GetName() string        { return p.name }
func (p Parallelogram) CalculateArea() float64 { return p.Base * p.Height }
