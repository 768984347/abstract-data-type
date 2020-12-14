package stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var expr = []string{"1", "+", "2", "*", "3", "-", "1"}
	num := CalExpression(expr)
	fmt.Println("express:", expr)
	fmt.Println("res:", num)
	if num != "6" {
		t.Errorf("res = %s,need %s", num, "6")
	}
	var expr2 = []string{"1", "/", "2"}
	num = CalExpression(expr2)
	fmt.Println("express:", expr2)
	fmt.Println("res:", num)
	if num != "0.5" {
		t.Errorf("res = %s,need %s", num, "0.5")
	}
	var expr3 = []string{"1", "/", "2", "*", "7", "-", "8", "*", "2", "+", "5"}
	num = CalExpression(expr3)
	fmt.Println("express:", expr3)
	fmt.Println("res:", num)
	if num != "-7.5" {
		t.Errorf("res = %s,need %s", num, "-7.5")
	}
}
