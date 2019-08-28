package implement_strstr

/**
@题目
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
@实现原理
将2个字符串各自进行相叠(如 aaacc => ['a','c'] [3,2]),相叠数量存储于另一个数组,再对匹配字符串进行遍历同时比对相叠数量(能够最大程度上减少匹配的次数提高代码效率)
需要注意边缘情况例如needle相叠后只有1位,最后一位只需要needle串相叠数量小于haystack串即可
@method strStr
@author pxb
@leetcode https://leetcode-cn.com/problems/implement-strstr/
*/
func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}
	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}
	haystackSimpleArr := []byte{haystack[0]}
	haystackSimpleNumArr := []int{1}
	haystackCurrIndex := 0
	needleSimpleArr := []byte{needle[0]}
	needleSimpleNumArr := []int{1}
	needleCurrIndex := 0

	for i := 1; i < len(haystack); i++ {
		if haystackSimpleArr[haystackCurrIndex] == haystack[i] {
			haystackSimpleNumArr[haystackCurrIndex]++
		} else {
			haystackSimpleArr = append(haystackSimpleArr, haystack[i])
			haystackSimpleNumArr = append(haystackSimpleNumArr, 1)
			haystackCurrIndex++
		}
	}

	for i := 1; i < len(needle); i++ {
		if needleSimpleArr[needleCurrIndex] == needle[i] {
			needleSimpleNumArr[needleCurrIndex]++
		} else {
			needleSimpleArr = append(needleSimpleArr, needle[i])
			needleSimpleNumArr = append(needleSimpleNumArr, 1)
			needleCurrIndex++
		}
	}
	for i := 0; i < len(haystackSimpleArr); i++ {
		//如果剩余的串长度已经小于要匹配的串则不需要匹配
		if len(haystackSimpleArr) - i < len(needleSimpleArr) {
			break
		}
		nextIndex := 0
		if needleSimpleArr[0] == haystackSimpleArr[i] {
			//满足当前字符长度继续匹配
			if needleSimpleNumArr[0] <= haystackSimpleNumArr[i] {
				j := 0
				if len(needleSimpleArr) == 1 {
					//当匹配字符串只有1位时只需要获取最前面的index
					if (needleSimpleArr[j] == haystackSimpleArr[i + j]) && (needleSimpleNumArr[j] <= haystackSimpleNumArr[i + j]) {
						return i
					}
				} else {
					j = 1
					for ; j < len(needleSimpleArr) - 1; j++ {
						if (needleSimpleArr[j] == haystackSimpleArr[i + j]) && (needleSimpleNumArr[j] == haystackSimpleNumArr[i + j]) {
							if needleSimpleArr[j] == needleSimpleArr[0] && nextIndex == 0 {
								nextIndex = i + j
							}
						} else {
							break
						}
					}
					//如果前面全部匹配判断最后一位字符是否匹配
					if j == len(needleSimpleArr) - 1 {
						if (needleSimpleArr[j] == haystackSimpleArr[i + j]) && (needleSimpleNumArr[j] <= haystackSimpleNumArr[i + j]) {
							//把前面相叠的长度相加
							returnNum := i
							for k := i - 1; k >= 0; k-- {
								if haystackSimpleNumArr[k] > 1 {
									returnNum += haystackSimpleNumArr[k] - 1
								}
							}
							return returnNum + haystackSimpleNumArr[i] - needleSimpleNumArr[0]
						}
					}
				}
				if nextIndex > 0 {
					i = nextIndex - 1
				} else {
					i += j - 1
				}
			}
		}
	}
	return -1
}