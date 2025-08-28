package main

import (
	"calculatorS/shapes"
	"fmt"

)

func main() {
	figures := []shapes.Shape{
		shapes.NewRectangle(10, 5),
		shapes.NewSquare(4),
		shapes.NewTriangle(6, 3),
		shapes.NewParallelogram(7, 3),
		shapes.NewCircle(3),
	}

	for _, f := range figures {
		fmt.Printf("Фигура [%s] имеет площадь %.2f\n", f.GetName(), f.CalculateArea())
	}
}
