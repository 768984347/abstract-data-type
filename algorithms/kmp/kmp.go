package kmp

/**
串的模式匹配-KMP算法
*/

/**
获取待匹配字符串模式数组(求字符串中每一位的最大重复子串长度的数组)
*/
func getNext(str string) []int {
	i := 0
	j := -1
	next := make([]int, len(str))
	next[0] = -1
	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

/**
KMP匹配算法
@param haystack 在该字符串中查找子串
@param needle 要匹配的子串
与BF匹配算法的区别是变量i不需要回溯,只要移动变量j
*/
func KMP(haystack string, needle string) int {
	next := getNext(needle)
	var i, j int
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(needle) {
		return i - j
	}
	return -1
}
