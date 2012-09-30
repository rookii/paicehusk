package paicehusk

import (
	"regexp"
	"strconv"
	"strings"
)

// A representation of a stemming rule

// A representation of a stemming rule

// A representation of a stemming rule

// A representation of a stemming rule
type rule struct {

	// The suffix the rule is to act on
	suf string

	// True if the stem is required intact for the rule to operate
	intact bool

	// Number of letters to strip off the stem
	num int

	// A suffix to append to the stem
	app string

	// True if the stem should be stemmed further
	cont bool
}

// RuleTable stores rules based on the final letter of the suffix they act on
type RuleTable struct {
	Table map[string][]*rule
}

// NewRuleTable returns a new RuleTable instance
func NewRuleTable(f []string) (table *RuleTable) {
	table = &RuleTable{Table: make(map[string][]*rule)}
	for _, s := range f {
		if r := ValidRule(s); r != "" {
			table.Table[r[:1]] = append(table.Table[r[:1]], ParseRule(r))
		}
	}
	return
}

// Validates a rule
func ValidRule(s string) (rule string) {
	reg := regexp.MustCompile("[a-zA-Z]*\\*?[0-9][a-zA-z]*[.>]")

	// Find the first instance of a rule in the provided string
	rule = reg.FindString(s)
	return
}

// Parses a rule in the form:
// |suffix|intact flag|number to strip|Append|Continue flag
//
// Eg, a rule: ht*2. Means if the stem is still intact, strip the
// suffix th and make no further attempts to stem the word.
//
// Rule nois4j> Means strip the sion suffix, append a j and check
// for any more rules to follow
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

// Stem a string, word, based on the rules in *RuleTable r, by following
// the algorithm described at:
// http://www.comp.lancs.ac.uk/computing/research/stemming/Links/paice.htm
func Stem(word string, r *RuleTable) string {
	stem := word

	// Intact Flag
	intact := true

	// If the stem is less than 3 chars, nothing to do, return
	if len(stem) < 3 {
		return stem
	}

	// Lookup the map to see if a rule is available for the
	// given stems last letter
	// A match was found
	if rules, ok := r.Table[stem[len(stem)-1:]]; ok {
		// Loop through the applicable rules
		for _, rule := range rules {
			// Don't bother if the length of the rule is greater than
			// the Stem
			if len(stem) > len(rule.suf) {
				// Check the rule matches
				if strings.HasSuffix(stem, Reverse(rule.suf)) {

					// If the strip count (rule.num) is set to 0 the stem
					// is protected and should be left alone
					if rule.num == 0 {
						break
					}

					// Apply the stem unless the intact flag is set and the
					// stem has been operated on allready
					if !((rule.intact == true) && (intact == false)) {
						// Check that the result of the rule is valid, otherwise
						// do nothing
						if s := stem[:len(stem)-rule.num]; ValidStem(s + rule.app) {
							stem = s + rule.app

							// Set the intact flag
							intact = false

							// Set the continue flag based on the rule
							if !rule.cont {
								break
							}
						}
					}
				}
			}
		}

	}
	return stem
}

// Acceptability condition: if the stem begins with a vowel, then it
// must contain at least 2 letters, one of which must be a consonant
//
// If however, it begins with a consonant then it must contain three
// letters and at least one of these must be a vowel or 'y'
func ValidStem(word string) bool {
	// If there's no vowel left in the stem, stem is invalid
	if !HasVowel(word) {
		return false
	}

	// If the word has a vowel and is longer than 3 letters, stem is valid
	if len(word) > 3 {
		return true
	}

	// If the first letter is a vowel
	if Vowel(word, 0) {
		if len(word) > 1 {
			// The Second letter must be a consonant
			if Consonant(word, 1) {
				return true
			}
		} else {
			return false
		}

	} else {
		// If the first letter is a consonant
		// The stem must contain 3 letters, one of which we allready know
		// to be a vowel
		if len(word) > 2 {
			return true
		}
	}
	return false
}

// Returns true if letter at offset is a consonant
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

// Returns true if letter at offset is a vowel
func Vowel(word string, offset int) bool {
	return !Consonant(word, offset)
}

// Returns true if the word contains a vowel
func HasVowel(word string) bool {
	for i := 0; i < len(word); i++ {
		if Vowel(word, i) {
			return true
		}
	}
	return false
}

// Reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
