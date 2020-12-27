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
		if j == -1 || str[i] == str[j] { //当前位已经没有可能重复的子串或者匹配字符相同时,向前移动2个指针
			i++
			j++
			next[i] = j
		} else {
			j = next[j] //当2个字符不相同时，回溯j
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
		if j == -1 || haystack[i] == needle[j] { //字符串首位都不匹配(j=-1)或者2个字符相同时，向前移动2个指针
			i++
			j++
		} else {
			j = next[j] //j赋值为当前位不重复的首位，避免不必要的匹配
		}
	}
	if j >= len(needle) {
		return i - j
	}
	return -1
}
