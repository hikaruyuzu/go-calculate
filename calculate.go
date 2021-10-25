package gocalculate

import (
	"fmt"
	"strconv"
)

type CalculateArea interface {
	WakeUpArea() int
}

type Rectangle struct {
	Length int
	Width  int
}

type Sequare struct {
	Side int
}

type Triangle struct {
	Base   int
	Height int
}

type Parallelogram struct {
	Base   int
	Height int
}

func (rectangle Rectangle) WakeUpArea() int {
	result := 2 * rectangle.Length * rectangle.Width
	return result
}

func (sequare Sequare) WakeUpArea() int {
	result := sequare.Side * sequare.Side
	return result
}

func (triangle Triangle) WakeUpArea() int {
	result := triangle.Base * triangle.Height / 2
	return result
}

func (parallelogram Parallelogram) WakeUpArea() int {
	result := parallelogram.Base * parallelogram.Height
	return result
}

func GetArea(area CalculateArea) string {
	result := area.WakeUpArea()
	resultStr := strconv.FormatInt(int64(result), 10)
	finalResult := fmt.Sprintf("Result: %s \n", resultStr)
	return finalResult
}
