package longest_common_prefix

/**
@题目 (最长公共前缀)
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

@实现原理
首先找出最小的元素,将最小元素的长度作为外层循环的长度来减少循环次数,当一个字符不匹配时直接结束循环
需要注意的点是当参数为空数组的边缘场景
@method strStr
@author pxb
@leetcode https://leetcode-cn.com/problems/longest-common-prefix
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lowerIndex := 0
	var returnStr []byte

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[lowerIndex]) {
			lowerIndex = i
		}
	}

	for i := 0; i < len(strs[lowerIndex]); i++ {
		isAppend := true
		for j := 0; j < len(strs); j++ {
			if j != lowerIndex {
				if strs[lowerIndex][i] != strs[j][i] {
					isAppend = false
					break
				}
			}
		}
		if isAppend {
			returnStr = append(returnStr, strs[lowerIndex][i])
		} else {
			break
		}
	}
	return string(returnStr)
}