package paicehusk

import (
	"strings"
	"testing"
)

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("Error: %v classified wrong. Expected %v, actual %v", msg, expected, actual)
	}
}

func TestConsonant(t *testing.T) {
	word := "THEY"
	compare(t, true, Consonant(word, 0), "T")
	compare(t, true, Consonant(word, 1), "H")
	compare(t, false, Consonant(word, 2), "E")
	compare(t, true, Consonant(word, 3), "Y")
	word = "Yellow"
	compare(t, true, Consonant(word, 0), "Yoke")
	word = "synergy"
	compare(t, true, Consonant(word, 0), "s")
	compare(t, false, Consonant(word, 1), "y")
	compare(t, true, Consonant(word, 2), "n")
	compare(t, false, Consonant(word, 3), "e")
	compare(t, true, Consonant(word, 4), "r")
	compare(t, true, Consonant(word, 5), "g")
	compare(t, false, Consonant(word, 6), "y")
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

func TestReverse(t *testing.T) {
	str := "Hello"
	expected := "olleH"
	if r := Reverse(str); r != expected {
		t.Errorf("Error: should be %v, got %v", expected, r)
	}
	str = "Here's a more complicated string to reverse."
	expected = ".esrever ot gnirts detacilpmoc erom a s'ereH"
	if r := Reverse(str); r != expected {
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
	if ok := ValidStem(test); ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "fire"
	if ok := ValidStem(test); !ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "aa"
	if ok := ValidStem(test); ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "ab"
	if ok := ValidStem(test); !ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "a"
	if ok := ValidStem(test); ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "ba"
	if ok := ValidStem(test); ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, false, ok)
	}
	test = "baa"
	if ok := ValidStem(test); !ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, true, ok)
	}
	test = "bba"
	if ok := ValidStem(test); !ok {
		t.Errorf("Error: ValidStem(%v) should be %v, got %v", test, true, ok)
	}
}

func TestStem(t *testing.T) {
	rSlice := strings.Split(ruleLiteral, "\n")
	table := NewRuleTable(rSlice)
	if len(table.Table) != 21 {
		t.Errorf("Error: len(table.Table) should be %v, got %v", 21, len(table.Table))
	}

	// To short
	expect := "at"
	if test := Stem("at", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// No rule match (No 'k' rules exist)
	expect = "rack"
	if test := Stem("rack", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// No rule match ('N' rules exist but no 'n', or 'no' rule)
	expect = "aaron"
	if test := Stem("aaron", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem has no vowels
	expect = "splat"
	if test := Stem("splat", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with consonat, and only has 2 letters
	expect = "doat"
	if test := Stem("doat", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with vowel, has 1 letters
	expect = "eat"
	if test := Stem("eat", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Resulting stem starts with vowel, has 2 letters
	expect = "ik"
	if test := Stem("ikat", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check protect rule
	expect = "foreseen"
	if test := Stem("foreseen", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check intact rule
	expect = "Aria"
	if test := Stem("Ariaan", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check replace rule
	expect = "explod"
	if test := Stem("explosion", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

	// Check partial replacement
	expect = "comply"
	if test := Stem("complicate", table); test != expect {
		t.Errorf("Error: expected %v, got %v", expect, test)
	}

}

var rule1 string = "ai*2."
var rule2 string = "lib3j>"
var rule3 string = "ab*2 ."
var rule4 string = "fire"
var rule5 string = "asfa __ falkjlk ?!@|.."
var rule6 string = "There's a rule here somewhere: afab*4fla>"

var ruleLiteral = `
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
