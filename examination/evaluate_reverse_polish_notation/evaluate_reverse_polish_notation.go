package evaluate_reverse_polish_notation

import "strconv"

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