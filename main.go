package main

import "fmt"

// Calculator interface with Add and Subtract methods
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

// Calculator struct
type calculator struct{}

// Add method
func (c calculator) Add(a, b int) int {
	return a + b
}

// Subtract method
func (c calculator) Subtract(a, b int) int {
	return a - b
}

func main() {
	var calc Calculator = calculator{}
	fmt.Println("10 + 5 =", calc.Add(10, 5))
	fmt.Println("10 - 5 =", calc.Subtract(10, 5))
}
