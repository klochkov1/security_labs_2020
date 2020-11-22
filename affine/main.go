package main

import (
	"flag"
	"fmt"
	"math/big"

	uaalpha "github.com/klochkov1/security_labs_2020/uaalpha"
)

func pmod(x, d int) int {
	m := x % d
	if x < 0 && d < 0 {
		m -= d
	}
	if x < 0 && d > 0 {
		m += d
	}
	return m
}

func getAffineCoder(a, b int) map[rune]rune {
	coder := make(map[rune]rune)
	alphaRunes := []rune(uaalpha.Alphabet)
	alphaRunes = alphaRunes[34:]
	m := len(alphaRunes)
	var orig, coded rune
	for i := 0; i < m; i++ {
		orig, coded = alphaRunes[i], alphaRunes[((a*i+b)%m)]
		coder[orig] = coded
		// fmt.Printf("%q : %q\n", string(orig), string(coded))
	}
	return coder
}

func getAffineDecoder(a, b int) map[rune]rune {
	decoder := make(map[rune]rune)
	alphaRunes := []rune(uaalpha.Alphabet)
	alphaRunes = alphaRunes[34:]
	m := len(alphaRunes)
	k := new(big.Int).ModInverse(big.NewInt(int64(a)), big.NewInt(int64(m)))
	mi := int(k.Int64())
	var orig, coded rune
	for i := 0; i < m; i++ {
		orig, coded = alphaRunes[i], alphaRunes[pmod(mi*(i-b), m)]
		decoder[orig] = coded
		// fmt.Printf("%q : %q\n", string(orig), string(coded))
	}
	return decoder
}

func codeRunes(coder map[rune]rune, orig []rune) []rune {
	coded := make([]rune, 0, len(orig))
	for _, v := range orig {
		coded = append(coded, codeRune(coder, v))
	}
	return coded
}

func codeRune(coder map[rune]rune, orig rune) rune {
	if k, ok := coder[orig]; ok {
		orig = k
	}
	return orig
}

func decodeRunes(decoder map[rune]rune, coded []rune) []rune {
	orig := make([]rune, 0, len(coded))
	for _, v := range coded {
		orig = append(orig, decodeRune(decoder, v))
		// fmt.Printf("%q : %q\n", string(v), string(decodeRune(decoder, v)))
	}
	return orig
}

func decodeRune(decoder map[rune]rune, orig rune) rune {
	if k, ok := decoder[orig]; ok {
		orig = k
	}
	return orig
}

func main() {
	var a, b int
	flag.IntVar(&a, "a", 1, "number of letters sequence to count in file")
	flag.IntVar(&b, "b", 1, "number of letters sequence to count in file")
	flag.Parse()
	str := "Привіт, як справи?"
	fmt.Println(str)
	strRunes := []rune(str)
	coder := getAffineCoder(a, b)
	decoder := getAffineDecoder(a, b)
	coded := codeRunes(coder, strRunes)
	fmt.Println(string(coded))
	fmt.Println(string(decodeRunes(decoder, coded)))
}
