package evaluate_reverse_polish_notation

import "strconv"

/**
逆波兰表达式求值
根据逆波兰表示法，求表达式的值。
有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
@实现原理
利用栈的后进先出原理,在循环中如果是运算符则将栈顶2个元素进行计算再入栈,否则直接将数字入栈
@author pxb
@method evaluate_reverse_polish_notation
@leetcode https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
*/
func evalRPN(tokens []string) int {
	var stackStr []string
	var res int64

	if len(tokens) < 2 {
		res, _ = strconv.ParseInt(tokens[0], 0, 0)
		return int(res)
	}

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			num1, _ := strconv.ParseInt(stackStr[len(stackStr) - 1], 0, 0)
			num2, _ := strconv.ParseInt(stackStr[len(stackStr) - 2], 0, 0)
			res = num2 + num1
			stackStr = stackStr[:len(stackStr) - 2]
			stackStr = append(stackStr, strconv.FormatInt(res, 10))
			break
		case "-":
			num1, _ := strconv.ParseInt(stackStr[len(stackStr) - 1], 0, 0)
			num2, _ := strconv.ParseInt(stackStr[len(stackStr) - 2], 0, 0)
			res = num2 - num1
			stackStr = stackStr[:len(stackStr) - 2]
			stackStr = append(stackStr, strconv.FormatInt(res, 10))
			break
		case "*":
			num1, _ := strconv.ParseInt(stackStr[len(stackStr) - 1], 0, 0)
			num2, _ := strconv.ParseInt(stackStr[len(stackStr) - 2], 0, 0)
			res = num2 * num1
			stackStr = stackStr[:len(stackStr) - 2]
			stackStr = append(stackStr, strconv.FormatInt(res, 10))
			break
		case "/":
			num1, _ := strconv.ParseInt(stackStr[len(stackStr) - 1], 0, 0)
			num2, _ := strconv.ParseInt(stackStr[len(stackStr) - 2], 0, 0)
			res = num2 / num1
			stackStr = stackStr[:len(stackStr) - 2]
			stackStr = append(stackStr, strconv.FormatInt(res, 10))
			break
		default:
			stackStr = append(stackStr, tokens[i])
		}
	}

	return int(res)
}