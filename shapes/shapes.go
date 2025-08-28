package shapes

// Общий интерфейс для всех фигур
type Shape interface {
	GetName() string
	CalculateArea() float64
}
