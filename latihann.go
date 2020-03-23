package main

import "fmt"

func main() {
	hasi, _, _, _, _ := Kal(16, 54, 23, 12, 130)
	fmt.Println("Hasil operasi hitung 16 x 54 + 23 / 12 - 130 adalah ", hasi)
}

func Kal(a int, b int, c int, d int, e int) (int, int, int, int, int) {
	hasil := ((a * b) + (c / d) - e)
	return hasil, 0, 0, 0, 0

}
