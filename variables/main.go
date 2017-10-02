package main

//import "fmt"

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

}
