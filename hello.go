package main

import "fmt"

func main() {
	mul, min := calc(5, 2)
	fmt.Println("mul:", mul, "min:", min)
}

func calc(a int, b int) (int, int) {
	multiply := a * b
	minus := a - b

	return multiply, minus
}
