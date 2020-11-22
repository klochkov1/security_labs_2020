package uaalpha

import (
	"testing"
)

func TestUaStringLess(t *testing.T) {
	if !uaStringLess("абї", "абц") {
		t.Errorf("Incorrect:  %q  <  %q", "абї", "абц")
	}
	if !uaStringLess(" ", "а") {
		t.Errorf("Incorrect:  %q  <  %q", " ", "a")
	}
	if !uaStringLess("А", "а") {
		t.Errorf("Incorrect:  %q  <  %q", "a", "A")
	}
	alphabet := []rune(Alphabet)
	for i := 1; i < len(alphabet); i++ {
		if uaStringLess(string(alphabet[i]), string(alphabet[i-1])) {
			t.Errorf("Incorrect: %q < %q", string(alphabet[i]), string(alphabet[i-1]))
		}
	}
	for i := 1; i < len(alphabet); i++ {
		if !uaStringLess(string(alphabet[i-1]), string(alphabet[i])) {
			t.Errorf("Incorrect: %q > %q", string(alphabet[i]), string(alphabet[i-1]))
		}
	}
}
