package luckygen

import "testing"

func TestRandomInt(t *testing.T) {
	num1 := RandomInt()
	num2 := RandomInt()

	if num1 == num2 {
		t.Errorf("The random numbers were identical (High likelyhood of error)")
	}
}

func TestRandomRune(t *testing.T) {
	rune1 := RandomRune()
	rune2 := RandomRune()

	if rune1 == rune2 {
		t.Errorf("The random runes were identical (High likelyhood of error)")
	}

	possibleMatches := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	matched := false
	for _, letter := range possibleMatches {
		if rune1 == letter {
			matched = true
			break
		}
	}
	if !matched {
		t.Errorf("%q is not a letter...", rune1)
	}
}
