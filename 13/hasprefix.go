package main

import (
	"fmt"
	"strings"
)

func main() {

	zle := "golang.org"
	dobrze := "http://golang.org"
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println(strings.HasPrefix("", "http://"))

	if strings.HasPrefix(zle, "http://") != true {
		fmt.Println("ZLE")
	}

	if strings.HasPrefix(dobrze, "http://") != true {
		fmt.Println("DOBRZE")
	}
}
