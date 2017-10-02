package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println("array a1", a1, "len", len(a1))
	const size = 2
	var a2 [size * 2]bool
	fmt.Println("array a2", a2, "len", len(a2))

	a3 :=  [...]int{1, 2, 3}
	fmt.Println("array a3", a3, "len", len(a3))
	a3[2] = 12
	fmt.Println("array a3", a3, "len", len(a3))

	var aa [3][2]int
	aa[1][1] = 1

	fmt.Println("array aa", aa, "len", len(aa))
}