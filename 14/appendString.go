// Go program to illustrate
// how to concatenate strings

package main

import (
	"fmt"
)

func main() {
	// Creating and initializing strings
	// using var keyword
	var str1 string
	str1 = "Welcome!"

	var str2 string
	str2 = "Geeksforgeeks"

	// Concatenating strings
	// Using + operator
	fmt.Println("New String 1: ", str1+" "+str2)

	// Creating and intiializing strings
	// Using shorthand declaration
	str3 := "Geeks"
	str4 := "Geeks"

	// Concatenating strings
	// Using + operator
	result := str3 + "for" + str4

	fmt.Println("New string 2: ", result)

}
