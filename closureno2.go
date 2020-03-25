package main

import "fmt"

func anonymous_func() func(string) string {

	return func(msg string) string {
		return msg
	}
}

func main() {
	print_func := anonymous_func()
	fmt.Println(print_func("Hi there anonymous"))

}
