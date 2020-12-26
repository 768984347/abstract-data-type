package brute_force

/**
串的模式匹配 Brute-Force(BF)算法
查找字符串首次出现的位置(needle是否是haystack的子串)
@param haystack 在该字符串中查找子串
@param needle 要匹配的子串
*/
func BF(haystack string, needle string) int {
	var i, j int
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j >= len(needle) {
		return i - j
	}
	return -1
}
