package main

import (
	"fmt"
)

func main() {
	eleven := 11   // decimal
	two := "2"     // hex
	three := 3     // oct
	four := 4      // binary
	eight := 8     // binary
	prawda := true // true
	falsz := false // false
	rune := 'キ'    // Unicode, Katakana Letter Ki
	letterA := 'a'
	stringA := "Ala ma kota" // quoted string %q
	veryBigNumber := 1234567890123.1
	lilFloat := 1.1

	fmt.Printf("%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\n")
	fmt.Printf("%d\t%x\t%o\t%b\t%b\t%b\t%b\n", eleven, two, three, four, eight, 256, 1234)
	fmt.Printf("%f\t%g\t%e\n", 3.123456789012345678901234567890, 3.123456789012345678901234567890, 3.123456789012345678901234567890)
	fmt.Printf("%t\t%t\n", prawda, falsz)
	fmt.Printf("%c = %x\t%c = %x\n", rune, rune, letterA, letterA)
	fmt.Printf("%s\n", two)
	fmt.Printf("%q\t%q\n", stringA, rune)
	fmt.Printf("%v\t%v\t%v\t%v\t%v\n", eleven, two, 1023, rune, prawda)
	fmt.Printf("%x\n", 12461) // proof of katakana unicode
	fmt.Printf("%T\t%T\t%T\t%T\t%T\t%T\t%T\t%T\t%T\n", eleven, two, prawda, falsz, rune, letterA, stringA, veryBigNumber, lilFloat)
	fmt.Printf("%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\t%%\n")

}
