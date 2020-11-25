package main

import (
	"bufio"
	"flag"
	"fmt"
	uaalpha "github.com/klochkov1/security_labs_2020/uaalpha"
	"io"
	"log"
	"math/big"
	"os"
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

func encodeRune(coder map[rune]rune, orig rune) rune {
	if k, ok := coder[orig]; ok {
		orig = k
	}
	return orig
}

func decodeRune(decoder map[rune]rune, orig rune) rune {
	if k, ok := decoder[orig]; ok {
		orig = k
	}
	return orig
}

func processStdin(coder map[rune]rune, processor func(map[rune]rune, rune) rune) {
	reader := bufio.NewReader(os.Stdin)
	// coded, err := os.Create("test.txt")
	// defer coded.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// writer := bufio.NewWriter(coded)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Print(string(processor(coder, r)))
		// writer.WriteRune(codeRune(coder, r))
		// writer.Flush()
	}
}

func main() {
	var a, b int
	flag.IntVar(&a, "a", -1, "number of letters sequence to count in file")
	flag.IntVar(&b, "b", -1, "number of letters sequence to count in file")
	decode := flag.Bool("decode", false, "decode stdin and print to stdout")
	encode := flag.Bool("encode", false, "encode stdin and print to stdout")
	flag.Parse()

	if (*decode && *encode) || !(*decode || *encode) || (a == -1 && b == -1) {
		fmt.Println("a and b values must be provided and either decode or encode flags must be set!")
		flag.PrintDefaults()
	}
	var coder map[rune]rune
	if *encode {
		coder = getAffineCoder(a, b)
		processStdin(coder, encodeRune)
	} else if *decode {
		coder = getAffineDecoder(a, b)
		processStdin(coder, decodeRune)
	}
}
