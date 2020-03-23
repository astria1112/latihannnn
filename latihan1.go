package main

import "fmt"

func main() {
	mul, div, add, min, _, _ := Calc(12, 144, 2, 34, 2345, 109)
	fmt.Println("multiply=", mul)
	fmt.Println("divide=", div)
	fmt.Println("add=", add)
	fmt.Println("minus=", min)
}

func Calc(a int, b int, c int, d int, e int, f int) (int, int, int, int, int, int) {
	multiply := a * a
	divide := b / c
	add := d + a
	minus := e - f

	return multiply, divide, add, minus, 0, 0

}
