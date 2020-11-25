package main

import (
	"flag"
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func getPrimeInts(m int) []int {
	var compInts []int
	for i := 1; i < m; i++ {
		n := gcd(i, m)
		if n == 1 {
			compInts = append(compInts, i)
		}
	}
	return compInts
}

func main() {
	var x, y, m int
	flag.IntVar(&m, "m", -1, "alphabet length")
	flag.IntVar(&x, "x", -1, "number of letters in source file")
	flag.IntVar(&y, "y", -1, "number of letters in destination file")
	flag.Parse()
	for _, a := range getPrimeInts(m) {
		for b := 0; b < m; b++ {
			if ((x*a + b) % m) == y {
				fmt.Printf("%d:%d\n", a, b)
			}
		}
	}
}
