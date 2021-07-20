package main

import (
	"fmt"
	"os"
	"strconv"
)

func Add(a, b float64) float64 {
	return a + b
}

func Decrease(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func main() {
	args := os.Args[1:]
	a, _ := strconv.ParseFloat(args[0], 64)
	b, _ := strconv.ParseFloat(args[1], 64)

	fmt.Println(Add(a, b))
	fmt.Println(Decrease(a, b))
	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))

	fmt.Println("Hello Go")
}
