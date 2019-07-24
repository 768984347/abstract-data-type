package plus_one

/**
@题目
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

@实现原理
从数组尾遍历数组项 + 1如果 > 9 需要进一位该位直接置0 进行下一次循环,否则直接结束遍历
最后如果数组首位是0则需要向数组首位插入1
@method plusOne
@author pxb
@leetcode https://leetcode-cn.com/problems/plus-one
*/

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 0 {
		newArr := []int{1}
		digits = append(newArr, digits...)
	}
	return digits
}
