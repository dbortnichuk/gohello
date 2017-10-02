package main

import "fmt"

func main() {

	i := 5
	b := &i
	*b = 6
	println(i)

	sl := []int{0, 0, 0}
	Tt(sl)
	fmt.Println(sl)


}

func Tt(items []int){
	items[2] = 1
}


