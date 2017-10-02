package main

import "fmt"

const (
	c1     = 1
	c2 int = 2
)

var (
	m1 string = ""
	m2        = 11
)

func main() {
	//println("hello world")

	//---variables---

	var i int = 10
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1
	var unsignedInt uint = 100500
	var unsignedBigInt uint64 = 1<<64 - 1

	println("integers: ", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	var f1 float32 = 3.14
	var f2 float64 = 3.14

	println("floats: ", f1, f2)

	var b = true

	println("boolean: ", b)

	var str1 = "hello\n\t"
	var str2 string = "world"

	println("string: ", str1+str2)

	escaping := `hello\n\t`
	println("escaping: ", escaping)

	var rawBinary byte = '\x27'

	println("rawBinary: ", rawBinary)

	//meaningOfLife := 42

	var i1 uint = 12
	var i2 int = 16

	println("sum: ", int(i1)+i2)

	println("int to string: " + string(i2))

	z := 3 + 2i
	println("complex numbers: ", z)

	fName := "Dmytro"
	lName := "Bortnichuk"
	fullName := fName + " " + lName

	println("my name is: "+fullName+", of length: ", len(fullName))

	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool

	println("default vals: ", defaultInt, defaultFloat, defaultString, defaultBool)

	var v1, v2 string = "v1", "v2"

	println(v1, v2)

	/* multi line comment
	var(
		m1 string = ""
		m2 = 11
	)

	println(m1, m2)

	*/

	const i5 int = 13
	const i6 = 1 // if no type specified for const, type inferred in the place of usage
	var i4 uint = 13
	println(i4 + uint(i5))
	println(i4 + i6)


	//---incrementor---

	const (
		zero  = iota
		one   = iota
		_
		three
	)

	println("iota zero:", zero)
	println("iota one:", one)
	println("iota three:", three)

	const (
		_  = iota //skip 0
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		//ZB
	)

	println("KB: ", KB)
	println("MB: ", MB)
	println("GB: ", GB)
	println("TB: ", TB)
	println("PB: ", PB)
	println("EB: ", EB)
	//println("ZB: ", ZB)

	const (
		flag1 = 1 << iota
		flag2
		flag3
	)

	println("flag1: ", flag1)
	println("flag2: ", flag2)
	println("flag3: ", flag3)
	println("flags: ", flag1 & flag2 & flag3)


	// --arrays---

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

	//---slices---

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




	return

}
