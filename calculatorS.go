package main

import (
	"fmt"
	"math"
)

// Прямоугольник
type rectangle struct {
	name   string
	length int
	width  int
}

func newRectangle(length, width int) rectangle {
	return rectangle{name: "Прямоугольник", length: length, width: width}
}

func (r rectangle) getName() string {
	return r.name
}

func (r rectangle) calculateArea() float64 {
	return float64(r.length * r.width)
}

// Квадрат
type square struct {
	name string
	side int
}

func newSquare(side int) square {
	return square{name: "Квадрат", side: side}
}

func (s square) getName() string {
	return s.name
}

func (s square) calculateArea() float64 {
	return float64(s.side * s.side)
}

// Треугольник
type triangle struct {
	name   string
	base   int
	height int
}

func newTriangle(base, height int) triangle {
	return triangle{name: "Треугольник", base: base, height: height}
}

func (t triangle) getName() string {
	return t.name
}

func (t triangle) calculateArea() float64 {
	return 0.5 * float64(t.base*t.height)
}

// Круг
type circle struct {
	name   string
	radius float64
}

func newCircle(radius float64) circle {
	return circle{name: "Круг", radius: radius}
}

func (c circle) getName() string {
	return c.name
}

func (c circle) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

// Параллелограмм
type parallelogram struct {
	name   string
	base   int
	height int
}

func newParallelogram(base, height int) parallelogram {
	return parallelogram{name: "Параллелограмм", base: base, height: height}
}

func (p parallelogram) getName() string {
	return p.name
}

func (p parallelogram) calculateArea() float64 {
	return float64(p.base * p.height)
}

// Главная функция
func main() {
	figures := []interface {
		getName() string
		calculateArea() float64
	}{
		newRectangle(10, 10),
		newSquare(5),
		newTriangle(6, 4),
		newCircle(3),
		newParallelogram(7, 3),
	}

	for _, fig := range figures {
		fmt.Printf("Фигура [%s] имеет площадь %.6f\n", fig.getName(), fig.calculateArea())
	}
}
