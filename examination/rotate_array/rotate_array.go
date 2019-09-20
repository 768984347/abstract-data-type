package rotate_array

import "math"

/**
@题目 (旋转数组)
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。

@实现原理
一个指针指向下一个要替换的元素的位置(secondIndex),一个指针指向当前要替换的元素位置(firstIndex),当2个指针相遇时将当前位++,下一个指针的位置重新计算
@method rotate
@author pxb
@leetcode https://leetcode-cn.com/problems/rotate-array
*/
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	if k >= len(nums) {
		k = k % len(nums)
	}
	if k == 0 {
		return
	}
	firstIndex := int(math.Abs(float64(len(nums) - k)))
	secondIndex := 0
	for i := 0; i < len(nums); i++ {
		if secondIndex == firstIndex {
			firstIndex++
			secondIndex = firstIndex
		} else {
			nums[secondIndex], nums[firstIndex] = nums[firstIndex], nums[secondIndex]
		}
		secondIndex = secondIndex + k
		if secondIndex > len(nums) - 1 {
			secondIndex = int(math.Abs(float64(len(nums) - secondIndex)))
		}
	}
}