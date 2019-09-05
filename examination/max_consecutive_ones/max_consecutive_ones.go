package max_consecutive_ones

/**
@题目 (最大连续1的个数)
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.

@实现原理
一个变量记录当前最大的连续个数,一个变量记录循环时最大的连续个数,如果不等于0就重置当前循环最大个数并且当个数大于maxLength时更新maxLength
需要注意的是只有一个元素时([1])的边缘场景
@method removeElement
@author pxb
@leetcode https://leetcode-cn.com/problems/max-consecutive-ones
*/
func findMaxConsecutiveOnes(nums []int) int {
	maxLength := 0
	nowLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nowLength++
		} else {
			if nowLength > maxLength {
				maxLength = nowLength
			}
			nowLength = 0
		}
	}
	if nowLength > maxLength {
		maxLength = nowLength
	}
	return maxLength
}