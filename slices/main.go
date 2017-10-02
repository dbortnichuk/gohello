package main

import "fmt"

func main() {

	var sl1 []int
	fmt.Println("empty slice:", sl1)

	sl1 = append(sl1, 100)
	fmt.Println("non empty slice:", sl1, "len", len(sl1), "cap", cap(sl1))

	sl2 :=  []int{1, 2, 3}
	fmt.Println("slice sl2", sl2, "len", len(sl2), "cap", cap(sl2))

	var sl3 []int
	sl3 = append(sl1, sl2...)
	fmt.Println("slice sl3", sl3, "len", len(sl3), "cap", cap(sl3))

	var slm [][]int
	slm = append(slm, sl2)
	fmt.Println("slice slm", slm, "len", len(slm), "cap", cap(slm))

	sl4 := make([]int, 10)
	fmt.Println("slice sl4", sl4, "len", len(sl4), "cap", cap(sl4))
	sl4 = append(sl4, 1)
	fmt.Println("slice sl4", sl4, "len", len(sl4), "cap", cap(sl4))

	sl5 := make([]int, 10, 12)
	fmt.Println("slice sl5", sl5, "len", len(sl5), "cap", cap(sl5))
	sl5 = append(sl5, []int{1, 2, 3}...)
	fmt.Println("slice sl5", sl5, "len", len(sl5), "cap", cap(sl5))

	sl5 = sl4
	sl5[5] = 999
	fmt.Println("slice sl5", sl5, "slice sl4", sl4)

	sl4 = append(sl4, 888)
	fmt.Println("slice sl5", sl5, "slice sl4", sl4)

	var sl6 []int

	copy(sl6, sl5) // incorrect slice copying

	fmt.Println("Slice6:", sl6)

	sl7 := make([]int, len(sl5), len(sl5))

	copy(sl7, sl5)

	fmt.Println("Slice7:", sl7)

	sl8 :=  []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println("sl8[2:4)", sl8[2:4])
	fmt.Println("sl8[2:]", sl8[2:])
	fmt.Println("sl8[:4)", sl8[:4])

	sl9 := append(sl8[:2], sl8[5:]...)

	fmt.Println("sl9:", sl9)

	a4 := [...]int{5, 6, 7}

	sl10 := a4[:]

	a4[1] = 10

	fmt.Println("sl10:", sl10)

}