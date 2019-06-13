package valid_parentheses

/**
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
@实现原理
循环中不是结尾符的字符入栈,如果是结尾符则与栈顶的字符进行匹配
@method isValid
@leetcode https://leetcode-cn.com/explore/learn/card/queue-stack/218/stack-last-in-first-out-data-structure/878/
*/

func isValid(s string) bool {
	var strSlice []int32
	for _, val := range s {
		if val == ')' {
			if len(strSlice) == 0 || strSlice[len(strSlice) - 1] != '(' {
				return false
			}
			strSlice = strSlice[:len(strSlice) - 1]
			continue
		} else if val == ']' {
			if len(strSlice) == 0 || strSlice[len(strSlice) - 1] != '[' {
				return false
			}
			strSlice = strSlice[:len(strSlice) - 1]
			continue
		} else if val == '}' {
			if len(strSlice) == 0 || strSlice[len(strSlice) - 1] != '{' {
				return false
			}
			strSlice = strSlice[:len(strSlice) - 1]
			continue
		} else {
			strSlice = append(strSlice, val)
		}
	}
	if len(strSlice) > 0 {
		return false
	}
	return true
}