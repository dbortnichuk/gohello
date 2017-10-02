package main

import "fmt"

func main() {

	var symbol rune = 'a'
	var autoSymbol = 'b'
	unicodeSymbol := 'Ї'
	unicodeSymbolByNumber := '\u2222'

	println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)
	fmt.Println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)

	str1 := "Привіт світ"
	fmt.Println("ua: ", str1, len(str1))

	for index, runeVal := range str1 {
		fmt.Printf("%#U at position %d\n", runeVal, index) //nicode symbol occupies two bytes
	}

	fmt.Println("symbol at index 2: ", str1[2])

	fmt.Println("---------")

	bin1 := []byte(str1)
	fmt.Println("binary ua: ", bin1, len(bin1))

	for idx, val := range bin1 {
		fmt.Printf("raw binary idx: %v, dec: %v, hex: %x\n", idx, val, val) //nicode symbol occupies two bytes
	}

}
