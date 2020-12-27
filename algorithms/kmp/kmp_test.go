package kmp

import (
	"testing"
)

func Test(t *testing.T) {
	pos := KMP("kacc", "kadd")
	if pos != -1 {
		t.Error("Pattern Fail[1]")
	}
	pos = KMP("aaaaaaaac", "ac")
	if pos != 7 {
		t.Error("Pattern Fail[2]")
	}
	pos = KMP("asdkclkj", "kc")
	if pos != 3 {
		t.Error("Pattern Fail[3]")
	}
	pos = KMP("aaab", "aaab")
	if pos != 0 {
		t.Error("Pattern Fail[4]")
	}
	pos = KMP("dasdksadsa", "dasdasdkajdlasd")
	if pos != -1 {
		t.Error("Pattern Fail[5]")
	}
	pos = KMP("llllllkcllllkcll", "kcl")
	if pos != 6 {
		t.Error("Pattern Fail[6]")
	}
}
