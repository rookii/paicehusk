package paicehusk

import (
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	suf    string
	intact bool
	num    int
	app    string
	cont   bool
}

type RuleTable struct {
	Table map[string][]*rule
}

func NewRuleTable(f []string) (table *RuleTable) {
	table = &RuleTable{Table: make(map[string][]*rule)}
	for _, s := range f {
		if r := ValidRule(s); r != "" {
			table.Table[r[:1]] = append(table.Table[r[:1]], ParseRule(r))
		}
	}
	return
}

func ValidRule(s string) (rule string) {
	reg := regexp.MustCompile("[a-zA-Z]*\\*?[0-9][a-zA-z]*[.>]")
	rule = reg.FindString(s)
	return
}

func ParseRule(s string) *rule {
	r := new(rule)
	suf := regexp.MustCompile("[a-zA-Z]+")
	intact := regexp.MustCompile("[*]")
	num := regexp.MustCompile("[0-9]")
	app := regexp.MustCompile("[0-9][a-zA-Z]+")
	r.suf = suf.FindString(s)
	if intact.FindString(s) == "" {
		r.intact = false
	} else {
		r.intact = true
	}
	if i, err := strconv.ParseInt(num.FindString(s), 0, 0); err != nil {
		panic(err)
	} else {
		r.num = int(i)
	}
	if append := app.FindString(s); len(append) > 0 {
		r.app = app.FindString(s)[1:]
	} else {
		r.app = ""
	}

	if s[len(s)-1:] == ">" {
		r.cont = true
	} else {
		r.cont = false
	}
	return r
}

func Stem(word string, r *RuleTable) string {
	cont := true
	stem := word

	if len(stem) < 3 {
		return stem
	}
	for cont {
		intact := true
		match := false
		for key := range r.Table {
			if stem[len(stem)-1:] == key {
				match = true
				cont = true
			} else {
				cont = false
			}
		}
		if match {
			rules := r.Table[stem[len(stem)-1:]]
			for _, rule := range rules {
				if len(stem) > len(rule.suf) {
					if strings.HasSuffix(stem, Reverse(rule.suf)) {
						if rule.num == 0 {
							break
						}
						if !((rule.intact == false) && (intact == false)) {
							// Stem the word
							if s := stem[:len(stem)-rule.num]; ValidStem(s) {
								intact = false
								cont = rule.cont
								stem = s
								// Apply any suffix
								if rule.app != "" {
									stem += rule.app
								}
							}
						}
					}
				}
			}
		}
	}
	return stem
}

func ValidStem(word string) bool {
	if len(word) > 3 {
		return true
	} else {
		if len(word) > 3 {
			if Vowel(word, 0) {
				if Consonant(word, 1) {
					return true
				}
			}
			if Consonant(word, 0) {
				if Vowel(word, 1) || Vowel(word, 2) {
					return true
				}
			}
		}
	}
	return false
}

func Consonant(word string, offset int) bool {
	switch word[offset] {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return false
	case 'Y', 'y':
		if offset == 0 {
			return true
		}
		return offset > 0 && !Consonant(word, offset-1)
	}
	return true
}

func Vowel(word string, offset int) bool {
	return !Consonant(word, offset)
}

func HasVowel(word string) bool {
	for i := 0; i < len(word); i++ {
		if Vowel(word, i) {
			return true
		}
	}
	return false
}

func FirstVowel(word string) int {
	for i := range word {
		if Vowel(word, i) {
			return i
		}
	}
	return 0
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
