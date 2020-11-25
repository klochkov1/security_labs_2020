package uaalpha

import (
	"testing"
)

func TestUaStringLess(t *testing.T) {
	if !UaStringLess("абї", "абц") {
		t.Errorf("Incorrect:  %q  <  %q", "абї", "абц")
	}
	if !UaStringLess(" ", "а") {
		t.Errorf("Incorrect:  %q  <  %q", " ", "a")
	}
	if !UaStringLess("А", "а") {
		t.Errorf("Incorrect:  %q  <  %q", "a", "A")
	}
	alphabet := []rune(Alphabet)
	for i := 1; i < len(alphabet); i++ {
		if UaStringLess(string(alphabet[i]), string(alphabet[i-1])) {
			t.Errorf("Incorrect: %q < %q", string(alphabet[i]), string(alphabet[i-1]))
		}
	}
	for i := 1; i < len(alphabet); i++ {
		if !UaStringLess(string(alphabet[i-1]), string(alphabet[i])) {
			t.Errorf("Incorrect: %q > %q", string(alphabet[i]), string(alphabet[i-1]))
		}
	}
}
