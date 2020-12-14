package stack

import (
	"abstract-data-type/stack"
	"strconv"
)

/**
表达式求值
@params expr 表达式数组 {"1","+","2","*","3"}
@return int  最后的值
*/
func CalExpression(expr []string) string {
	opv := stack.Init((len(expr) + 1) / 2) //操作数栈，存储数字1,2,3...
	opt := stack.Init((len(expr)+1)/2 - 1) //运算符号栈，存储运算符+,-,*,/...
	optMap := map[string]int{              //运算符优先级map 运算符 => 优先级
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
	for _, v := range expr {
		switch v {
		case "+", "-", "*", "/": //匹配运算符
			for {
				if opt.IsEmpty() {
					opt.Push(v)
					break
				}
				top := opt.Top().(string)
				if optMap[v] > optMap[top] { //运算符和栈顶运算符比较，如果优先级高，放入栈中
					opt.Push(v)
					break
				} else { //优先级低或相同，取栈顶2个数和运算符栈顶符号计算，计算后的数压入操作数栈，继续循环比较直到运算符栈空或者优先级高于栈顶的运算符
					popNum1, _ := opv.Pop()
					popNum2, _ := opv.Pop()
					optSymbol, _ := opt.Pop()
					res := cal(popNum2.(string), popNum1.(string), optSymbol.(string))
					opv.Push(res)
				}
			}
		default: //匹配数字
			opv.Push(v) //操作数直接入栈
		}
	}
	//执行到此处必定只剩2个操作数和1个运算符合，清空栈得到最后的结果
	popNum1, _ := opv.Pop()
	popNum2, _ := opv.Pop()
	optSymbol, _ := opt.Pop()
	res := cal(popNum2.(string), popNum1.(string), optSymbol.(string))
	return res
}

/**
加减乘除运算
*/
func cal(num1 string, num2 string, opt string) string {
	num1Float, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		panic(err)
	}
	num2Float, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		panic(err)
	}
	var res float64
	switch opt {
	case "+":
		res = num1Float + num2Float
	case "-":
		res = num1Float - num2Float
	case "*":
		res = num1Float * num2Float
	case "/":
		res = num1Float / num2Float
	}
	return strconv.FormatFloat(res, 'f', -1, 64)
}
