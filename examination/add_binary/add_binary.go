package add_binary

/**
@题目
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
@实现原理
只循环长度较长的那个字符串,当较短的字符串长度已经用完时只需要插入较长字符串的字符,通过记录是否自增的变量来判断当前位是否需要自增
@method addBinary
@author pxb
@leetcode https://leetcode-cn.com/problems/add-binary
*/
func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	var returnStr,tempStr []byte
	var loopStr, anotherStr string
	var anotherCurrIndex int
	var isIncr bool
	aLen := len(a)
	bLen := len(b)

	if aLen > bLen {
		loopStr = a
		anotherStr = b
		anotherCurrIndex = bLen - 1
	} else {
		loopStr = b
		anotherStr = a
		anotherCurrIndex = aLen - 1
	}

	for i := len(loopStr) - 1; i >= 0; i-- {
		var tempNum byte
		if anotherCurrIndex >= 0 {
			if loopStr[i] == '1' {
				if anotherStr[anotherCurrIndex] == '1' {
					if isIncr {
						tempNum = '1'
					} else {
						tempNum = '0'
					}
					isIncr = true
					anotherCurrIndex--
					tempStr = append(tempStr, tempNum)
					continue
				} else {
					tempNum = '1'
				}
			} else {
				if anotherStr[anotherCurrIndex] == '1' {
					tempNum = '1'
				} else {
					tempNum = '0'
				}
			}
			anotherCurrIndex--
		} else {
			tempNum = loopStr[i]
		}
		if isIncr {
			if tempNum == '1' {
				tempNum = '0'
			} else {
				tempNum = '1'
				isIncr = false
			}
		}
		tempStr = append(tempStr, tempNum)
	}
	if isIncr {
		tempStr = append(tempStr, '1')
	}
	for i := len(tempStr) - 1; i >= 0; i-- {
		returnStr = append(returnStr, tempStr[i])
	}
	return string(returnStr)
}