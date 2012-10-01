// Go implementation of the Paice/Husk Stemming algorithm:
// http://www.comp.lancs.ac.uk/computing/research/stemming/Links/paice.htm
package paicehusk

import (
	"regexp"
	"strconv"
	"strings"
)

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

var DefaultRules = NewRuleTable(strings.Split(defaultRules, "\n"))

// RuleTable stores rules based on the final letter of the suffix they act on
type RuleTable struct {
	Table map[string][]*rule
}

// NewRuleTable returns a new RuleTable instance
func NewRuleTable(f []string) (table *RuleTable) {
	table = &RuleTable{Table: make(map[string][]*rule)}
	for _, s := range f {
		if r, ok := ParseRule(s); ok {
			table.Table[r.suf[:1]] = append(table.Table[r.suf[:1]], r)
		}
	}
	return
}

// Validates a rule
func ValidRule(s string) (rule string, ok bool) {
	reg := regexp.MustCompile("[a-zA-Z]*\\*?[0-9][a-zA-z]*[.>]")
	ok = true
	// Find the first instance of a rule in the provided string
	rule = reg.FindString(s)
	if rule == "" {
		ok = false
	}
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
func ParseRule(s string) (r *rule, ok bool) {
	s, ok = ValidRule(s)
	if !ok {
		return nil, false
	}

	r = new(rule)
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
	return r, true
}

// Stem a string, word, based on the rules in *RuleTable r, by following
// the algorithm described at:
// http://www.comp.lancs.ac.uk/computing/research/stemming/Links/paice.htm
func (r *RuleTable) Stem(word string) string {
	stem := word
	current := word

	// Intact Flag
	intact := true

	// If the stem is less than 3 chars, nothing to do, return
	if len(stem) < 3 {
		return stem
	}

	// Main Control Loop
	cont := true
	for cont {
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
								cont = rule.cont
								current = s + rule.app

								// Set the intact flag
								intact = false

								// Set the continue flag based on the rule
								break
							}
						}
					}
				}
			}
			// No rule matched
			if current == stem {
				cont = false
			} else {
				// Set the new stem
				stem = current
			}
		} else {
			// No matching rule
			break
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

// Default Paice/Husk Rules
var defaultRules = `
ai*2.     { -ia > -   if intact }
a*1.      { -a > -    if intact }
bb1.      { -bb > -b   }
city3s.   { -ytic > -ys }
ci2>      { -ic > -    }
cn1t>     { -nc > -nt  }
dd1.      { -dd > -d   }
dei3y>    { -ied > -y  }
deec2ss.  { -ceed > -cess }
dee1.     { -eed > -ee }
de2>      { -ed > -    }
dooh4>    { -hood > -  }
e1>       { -e > -     }
feil1v.   { -lief > -liev }
fi2>      { -if > -    }
gni3>     { -ing > -   }
gai3y.    { -iag > -y  }
ga2>      { -ag > -    }
gg1.      { -gg > -g   }
ht*2.     { -th > -   if intact }
hsiug5ct. { -guish > -ct }
hsi3>     { -ish > -   }
i*1.      { -i > -    if intact }
i1y>      { -i > -y    }
ji1d.     { -ij > -id   --  see nois4j> & vis3j> }
juf1s.    { -fuj > -fus }
ju1d.     { -uj > -ud  }
jo1d.     { -oj > -od  }
jeh1r.    { -hej > -her }
jrev1t.   { -verj > -vert }
jsim2t.   { -misj > -mit }
jn1d.     { -nj > -nd  }
j1s.      { -j > -s    }
lbaifi6.  { -ifiabl > - }
lbai4y.   { -iabl > -y }
lba3>     { -abl > -   }
lbi3.     { -ibl > -   }
lib2l>    { -bil > -bl }
lc1.      { -cl > c    }
lufi4y.   { -iful > -y }
luf3>     { -ful > -   }
lu2.      { -ul > -    }
lai3>     { -ial > -   }
lau3>     { -ual > -   }
la2>      { -al > -    }
ll1.      { -ll > -l   }
mui3.     { -ium > -   }
mu*2.     { -um > -   if intact }
msi3>     { -ism > -   }
mm1.      { -mm > -m   }
nois4j>   { -sion > -j }
noix4ct.  { -xion > -ct }
noi3>     { -ion > -   }
nai3>     { -ian > -   }
na2>      { -an > -    }
nee0.     { protect  -een }
ne2>      { -en > -    }
nn1.      { -nn > -n   }
pihs4>    { -ship > -  }
pp1.      { -pp > -p   }
re2>      { -er > -    }
rae0.     { protect  -ear }
ra2.      { -ar > -    }
ro2>      { -or > -    }
ru2>      { -ur > -    }
rr1.      { -rr > -r   }
rt1>      { -tr > -t   }
rei3y>    { -ier > -y  }
sei3y>    { -ies > -y  }
sis2.     { -sis > -s  }
si2>      { -is > -    }
ssen4>    { -ness > -  }
ss0.      { protect  -ss }
suo3>     { -ous > -   }
su*2.     { -us > -   if intact }
s*1>      { -s > -    if intact }
s0.       { -s > -s    }
tacilp4y. { -plicat > -ply }
ta2>      { -at > -    }
tnem4>    { -ment > -  }
tne3>     { -ent > -   }
tna3>     { -ant > -   }
tpir2b.   { -ript > -rib }
tpro2b.   { -orpt > -orb }
tcud1.    { -duct > -duc }
tpmus2.   { -sumpt > -sum }
tpec2iv.  { -cept > -ceiv }
tulo2v.   { -olut > -olv }
tsis0.    { protect  -sist }
tsi3>     { -ist > -   }
tt1.      { -tt > -t   }
uqi3.     { -iqu > -   } 
ugo1.     { -ogu > -og }
vis3j>    { -siv > -j  }
vie0.     { protect  -eiv }
vi2>      { -iv > -    }
ylb1>     { -bly > -bl }
yli3y>    { -ily > -y  }
ylp0.     { protect  -ply }
yl2>      { -ly > -    }
ygo1.     { -ogy > -og }
yhp1.     { -phy > -ph }
ymo1.     { -omy > -om }
ypo1.     { -opy > -op }
yti3>     { -ity > -   }
yte3>     { -ety > -   }
ytl2.     { -lty > -l  }
yrtsi5.   { -istry > - }
yra3>     { -ary > -   }
yro3>     { -ory > -   }
yfi3.     { -ify > -   }
ycn2t>    { -ncy > -nt }
yca3>     { -acy > -   }
zi2>      { -iz > -    }
zy1s.     { -yz > -ys  }
end0.
`
