package largest_number_at_least_twice_of_others

/**
@题目
在一个给定的数组nums中，总是存在一个最大元素 。
查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
如果是，则返回最大元素的索引，否则返回-1。

示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.
 
示例 2:

输入: nums = [1, 2, 3, 4]
输出: -1
解释: 4没有超过3的两倍大, 所以我们返回 -1.

提示:

nums 的长度范围在[1, 50].
每个 nums[i] 的整数范围在 [0, 99].
@实现原理
在循环数组的过程中比对大小,使用2个变量记住数组中最大一个数字的index(maxNumIndex)和第二大数字的index(secNumIndex)
最后只要最大的数字 >= 第二大数字 * 2 就能得出最大元素是否至少是数组中每个其他数字的两倍的数字
@author pxb
@method dominantIndex
@leetcode https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
 */
func dominantIndex(nums []int) int {
	length := len(nums)
	if length > 0 {
		if length == 1 {
			return 0
		}
		maxNumIndex := 0
		secNumIndex := 0
		if nums[0] > nums[1] {
			secNumIndex = 1
		} else {
			maxNumIndex = 1
		}
		for i := 2; i < length; i++ {
			if nums[i] > nums[maxNumIndex] {
				secNumIndex = maxNumIndex
				maxNumIndex = i
			} else if nums[i] > nums[secNumIndex] {
				secNumIndex = i
			}
		}
		if nums[maxNumIndex] >= (nums[secNumIndex] * 2) {
			return maxNumIndex
		}
	}
	return -1
}