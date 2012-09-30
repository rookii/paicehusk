package paicehusk

import "testing"

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("%v -- value differs. Expected %v, actual %v", msg, expected, actual)
	}
}

func TestConsonant(t *testing.T) {
	word := "BUOY"
	compare(t, true, Consonant(word, 0), "B")
	compare(t, false, Consonant(word, 1), "U")
	compare(t, false, Consonant(word, 2), "O")
	compare(t, true, Consonant(word, 3), "Y")
	word = "synergy"
	compare(t, true, Consonant(word, 0), "s")
	compare(t, false, Consonant(word, 1), "y")
	compare(t, true, Consonant(word, 2), "n")
	compare(t, false, Consonant(word, 3), "e")
	compare(t, true, Consonant(word, 4), "r")
	compare(t, true, Consonant(word, 5), "g")
	compare(t, false, Consonant(word, 6), "y")
	word = "Yoke"
	compare(t, true, Consonant(word, 0), "Yoke")
}

func TestVowel(t *testing.T) {
	word := "BUOY"
	compare(t, false, Vowel(word, 0), "B")
	compare(t, true, Vowel(word, 1), "U")
	compare(t, true, Vowel(word, 2), "O")
	compare(t, false, Vowel(word, 3), "Y")
	word = "synergy"
	compare(t, false, Vowel(word, 0), "s")
	compare(t, true, Vowel(word, 1), "y")
	compare(t, false, Vowel(word, 2), "n")
	compare(t, true, Vowel(word, 3), "e")
	compare(t, false, Vowel(word, 4), "r")
	compare(t, false, Vowel(word, 5), "g")
	compare(t, true, Vowel(word, 6), "y")
	word = "Yoke"
	compare(t, false, Vowel(word, 0), "Yoke")
}

func TestValidRule(t *testing.T) {
	if _, ok := ValidRule(rule1); !ok {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}
	if _, ok := ValidRule(rule2); !ok {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}
	if _, ok := ValidRule(rule3); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}
	if _, ok := ValidRule(rule4); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}
	if _, ok := ValidRule(rule5); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}
	if _, ok := ValidRule(rule6); !ok {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}
}

var rule1 string = "ai*2."
var rule2 string = "lib2l>"
var rule3 string = "ab*2 ."
var rule4 string = "fire"
var rule5 string = "asfa __ falkjlk ?!@|.."
var rule6 string = "There's a rule here somewhere: afab*4fla>"
