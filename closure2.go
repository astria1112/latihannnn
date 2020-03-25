package main

import "fmt"

func main() {
	anom_func := func(msg string) {
		fmt.Println(msg)
		return
	}
	anom_func("hello there anonymous")

}
