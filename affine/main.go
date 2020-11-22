package main

import (
	"fmt"
	uaAlpha "github.com/klochkov1/ua_letters_frequency"
)

func affine(a, b int) {
	for i := 0; i < len(uaAlpha.Alphabet); i++ {
		fmt.Printf("%v\n", []rune(uaAlpha.Alphabet))
	}
}

func main() {
	str := "Hello, world."
	_ = str
	result := ""
	fmt.Println(result)
	fmt.Println(uaAlpha.Alphabet)
}
