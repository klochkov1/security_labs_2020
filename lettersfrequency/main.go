package main

import (
	"bufio"
	"flag"
	"fmt"
	uaalpha "github.com/klochkov1/security_labs_2020/uaalpha"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type charsCount struct {
	Chars string
	Count int
}

func insertSorted(ccl []charsCount, chars string) []charsCount {
	i := 0
	for ; i < len(ccl); i++ {
		if ccl[i].Chars == chars {
			ccl[i].Count++
			return ccl
		} else if uaalpha.UaStringLess(chars, ccl[i].Chars) {
			break
		}
	}
	ccl = append(ccl, charsCount{})
	copy(ccl[i+1:], ccl[i:])
	ccl[i] = charsCount{chars, 1}
	return ccl
}

func countCharsInFile(file *os.File, sequenceLength int, allLower bool) []charsCount {
	ccl := []charsCount{}

	reader := bufio.NewReader(file)
	seq := ""
	combinations := sequenceLength

	for {
		for i := 0; i < sequenceLength; i++ {
			ch, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					combinations--
					if combinations == 0 {
						return ccl
					}
					file.Seek(int64(sequenceLength-combinations), 0)
					break
				} else {
					log.Fatal(err)
				}
			}
			_, ok := uaalpha.AlphaIndex[string(ch)]
			if ok {
				if allLower {
					ch = unicode.ToLower(ch)
				}
				seq += string(ch)
			} else {
				seq = ""
				break
			}
		}
		if seq != "" {
			ccl = insertSorted(ccl, seq)
			seq = ""
		}
	}
}

func maxCount(cc []charsCount) int {
	max := 0
	for i := 0; i < len(cc); i++ {
		if cc[i].Count > max {
			max = cc[i].Count
		}
	}
	return max
}

func printDiagram(cc []charsCount) {
	ttyWidth, err := terminal.Width()
	if err != nil {
		log.Println("Failed to get tty width, using 100 as default value.")
		ttyWidth = 100
	}
	maxCount := maxCount(cc)
	maxd := 0
	for d := maxCount; d != 0; d /= 10 {
		maxd++
	}
	maxWidth := int(ttyWidth) - maxd - len([]rune(cc[0].Chars)) - 5 // quotes and spaces
	n := (maxCount / maxWidth) + 1                                  // to round up division
	for i := 0; i < len(cc); i++ {
		p := strings.Repeat("#", cc[i].Count/n)
		fmt.Printf("%3q %5d %v\n", cc[i].Chars, cc[i].Count, p)
	}
}

func main() {
	var sequenceLength int
	flag.IntVar(&sequenceLength, "n", 1, "number of letters sequence to count in file")
	allLower := flag.Bool("i", false, "ignore case, all letters will be converted to lower case ")
	stdin := flag.Bool("s", true, "read from stdin")
	flag.Parse()

	files := flag.Args()

	for i := 0; i < len(files) || *stdin; i++ {
		var file *os.File
		var err error
		if len(files) == 0 && *stdin {
			log.Println("No file arguments provided, reading from standart input.")
			file = os.Stdin
		} else {
			*stdin = false
			file, err = os.Open(files[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		result := countCharsInFile(file, sequenceLength, *allLower)
		fmt.Printf("\n%v sorted alphabetically:\n\n", file.Name())
		printDiagram(result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Count > result[j].Count
		})
		fmt.Printf("\n\n%v sorted by frequency:\n\n", file.Name())
		printDiagram(result)
		defer file.Close()
	}
}
