package paicehusk

import (
	"testing"
)

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("Error: %v classified wrong. Expected %v, actual %v", msg, expected, actual)
	}
}

// Mostly checking for the Y special cases
func TestConsonant(t *testing.T) {
	var word []rune = []rune{'T', 'H', 'E', 'Y'}
	compare(t, true, consonant(word, 0), "T in THEY")
	compare(t, true, consonant(word, 1), "H in THEY")
	compare(t, false, consonant(word, 2), "E in THEY")
	compare(t, true, consonant(word, 3), "Y in THEY")
	word = []rune{'Y', 'O', 'K', 'E'}
	compare(t, true, consonant(word, 0), "Yoke")
	word = []rune{'s', 'y', 'n', 'e', 'r', 'g', 'y'}
	compare(t, true, consonant(word, 0), "s in synergy")
	compare(t, false, consonant(word, 1), "y in synergy")
	compare(t, true, consonant(word, 2), "n in synergy")
	compare(t, false, consonant(word, 3), "e in synergy")
	compare(t, true, consonant(word, 4), "r in synergy")
	compare(t, true, consonant(word, 5), "g in synergy")
	compare(t, false, consonant(word, 6), "y in synergy")
	// Unicode test (I hope?)
	word = []rune{'男', '孩', 'b', 'o', 'y'}
	compare(t, true, consonant(word, 2), "b in 男孩boy")
	compare(t, false, consonant(word, 3), "o in 男孩boy")
	compare(t, true, consonant(word, 4), "y in 男孩boy")
}

// Same again
func TestVowel(t *testing.T) {
	var word = []rune{'B', 'U', 'O', 'Y'}
	compare(t, false, vowel(word, 0), "B in BUOY")
	compare(t, true, vowel(word, 1), "U in BUOY")
	compare(t, true, vowel(word, 2), "O in BUOY")
	compare(t, false, vowel(word, 3), "Y in BUOY")
	word = []rune{'s', 'y', 'n', 'e', 'r', 'g', 'y'}
	compare(t, false, vowel(word, 0), "s in synergy")
	compare(t, true, vowel(word, 1), "y in synergy")
	compare(t, false, vowel(word, 2), "n in synergy")
	compare(t, true, vowel(word, 3), "e in synergy")
	compare(t, false, vowel(word, 4), "r in synergy")
	compare(t, false, vowel(word, 5), "g in synergy")
	compare(t, true, vowel(word, 6), "y in synergy")
	word = []rune{'Y', 'o', 'k', 'e'}
	compare(t, false, vowel(word, 0), "Yoke")
	// Unicode test
	word = []rune{'男', '孩', 'b', 'o', 'y'}
	compare(t, false, vowel(word, 2), "b in 男孩boy")
	compare(t, true, vowel(word, 3), "o in 男孩boy")
	compare(t, false, vowel(word, 4), "y in 男孩boy")
}

func TestReverse(t *testing.T) {
	str := "Hello"
	expected := "olleH"
	if r := reverse(str); r != expected {
		t.Errorf("Error: should be %v, got %v", expected, r)
	}
	str = "Here's a more complicated string to reverse."
	expected = ".esrever ot gnirts detacilpmoc erom a s'ereH"
	if r := reverse(str); r != expected {
		t.Errorf("Error: should be %v, got %v", expected, r)
	}
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

func TestParseRule(t *testing.T) {
	if r, ok := ParseRule(rule1); ok {
		if r.suf != "ai" {
			t.Errorf("Error: r.suf should be %v, got %v", "ai", r.suf)
		}
		if r.intact != true {
			t.Errorf("Error: r.intact should be %v, got %v", true, r.intact)
		}
		if r.num != 2 {
			t.Errorf("Error: r.num should be %v, got %v", "2", r.num)
		}
		if r.app != "" {
			t.Errorf("Error: r.app should be %v, got %v", "", r.app)
		}
		if r.cont != false {
			t.Errorf("Error: r.cont should be %v, got %v", false, r.cont)
		}
	} else {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}

	if r, ok := ParseRule(rule2); ok {
		if r.suf != "lib" {
			t.Errorf("Error: r.suf should be %v, got %v", "lib", r.suf)
		}
		if r.intact != false {
			t.Errorf("Error: r.intact should be %v, got %v", false, r.intact)
		}
		if r.num != 3 {
			t.Errorf("Error: r.num should be %v, got %v", "3", r.num)
		}
		if r.app != "j" {
			t.Errorf("Error: r.app should be %v, got %v", "j", r.app)
		}
		if r.cont != true {
			t.Errorf("Error: r.cont should be %v, got %v", true, r.cont)
		}
	} else {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}

	if _, ok := ParseRule(rule3); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}

	if _, ok := ParseRule(rule4); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}

	if _, ok := ParseRule(rule5); ok {
		t.Errorf("Error: Ok should be %v, got %v", false, ok)
	}

	if r, ok := ParseRule(rule6); ok {
		if r.suf != "afab" {
			t.Errorf("Error: r.suf should be %v, got %v", "afab", r.suf)
		}
		if r.intact != true {
			t.Errorf("Error: r.intact should be %v, got %v", true, r.intact)
		}
		if r.num != 4 {
			t.Errorf("Error: r.num should be %v, got %v", "4", r.num)
		}
		if r.app != "fla" {
			t.Errorf("Error: r.app should be %v, got %v", "fla", r.app)
		}
		if r.cont != true {
			t.Errorf("Error: r.cont should be %v, got %v", true, r.cont)
		}
	} else {
		t.Errorf("Error: Ok should be %v, got %v", true, ok)
	}
}

func TestNewRuleTable(t *testing.T) {
	f := []string{rule1, rule2, rule3, rule4, rule5, rule6}
	table := NewRuleTable(f)
	if len(table.Table) != 2 {
		t.Errorf("Error: len(table.Table) should be %v, got %v", 2, len(table.Table))
	}
	if len(table.Table["a"]) != 2 {
		t.Errorf("Error: len(table.Table[\"a\"]) should be %v, got %v", 2, len(table.Table))
	}
}

func TestValidStem(t *testing.T) {
	test := "xvzf"
	if ok := validStem(test); ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "fire"
	if ok := validStem(test); !ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "aa"
	if ok := validStem(test); ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "ab"
	if ok := validStem(test); !ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "a"
	if ok := validStem(test); ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "ba"
	if ok := validStem(test); ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "baa"
	if ok := validStem(test); !ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "bba"
	if ok := validStem(test); !ok {
		t.Errorf("Error: validStem(%v) should be %v, got %v", test, true, ok)
	}
}

func TestStem(t *testing.T) {
	// To short
	expect := "at"
	if test := DefaultRules.Stem("at"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// No rule match (No 'k' rules exist)
	expect = "rack"
	if test := DefaultRules.Stem("rack"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// No rule match ('N' rules exist but no 'n', or 'no' rule)
	expect = "aaron"
	if test := DefaultRules.Stem("aaron"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem has no vowels
	expect = "splat"
	if test := DefaultRules.Stem("splat"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with consonat, and only has 2 letters
	expect = "doat"
	if test := DefaultRules.Stem("doat"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with vowel, has 1 letters
	expect = "eat"
	if test := DefaultRules.Stem("eat"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with vowel, has 2 letters
	expect = "ik"
	if test := DefaultRules.Stem("ikat"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check protect rule
	expect = "foreseen"
	if test := DefaultRules.Stem("foreseen"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check intact rule
	expect = "Aria"
	if test := DefaultRules.Stem("Ariaan"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check replace rule
	expect = "explod"
	if test := DefaultRules.Stem("explosion"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check partial replacement
	expect = "comply"
	if test := DefaultRules.Stem("complicate"); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

}

// Test for the Parse
var rule1 string = "ai*2."
var rule2 string = "lib3j>"
var rule3 string = "ab*2 ."
var rule4 string = "fire"
var rule5 string = "asfa __ falkjlk ?!@|.."
var rule6 string = "There's a rule here somewhere: afab*4fla>"
