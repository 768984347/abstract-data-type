package brute_force

import (
	"testing"
)

func Test(t *testing.T) {
	str1 := "aaaaaab"
	str2 := "aaab"
	pos := BF(str1, str2)
	if pos != 3 {
		t.Error("Pattern Fail[1]")
	}
	str1 = "aaab"
	str2 = "aaab"
	pos = BF(str1, str2)
	if pos != 0 {
		t.Error("Pattern Fail[2]")
	}
	str1 = "aaab"
	str2 = "ac"
	pos = BF(str1, str2)
	if pos != -1 {
		t.Error("Pattern Fail[3]")
	}
	str1 = "hhkcgg"
	str2 = "k"
	pos = BF(str1, str2)
	if pos != 2 {
		t.Error("Pattern Fail[3]")
	}
	str1 = "ddkcbb"
	str2 = "kc"
	pos = BF(str1, str2)
	if pos != 2 {
		t.Error("Pattern Fail[3]")
	}
}
