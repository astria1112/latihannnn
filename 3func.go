package main

import "fmt"

func main() {
	b := b()
	slice := []int{45, 23, 56, 89, 900, 26, 78}
	cc, dd := b(slice...)
	fmt.Println("genap:", cc)
	fmt.Println("ganjil:", dd)
}

func b() func(...int) ([]int, []int) {
	return func(a ...int) ([]int, []int) {
		var genap = []int{}
		var ganjil = []int{}
		for i := 0; i < len(a); i++ {
			if a[i]%2 == 0 {
				genap = append(genap, a[i])
			} else {
				ganjil = append(ganjil, a[i])
			}
		}
		return genap, ganjil
	}

}
