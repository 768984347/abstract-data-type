package minimum_size_subarray_sum

/**
@题目 (长度最小的子数组)
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例: 

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

@实现原理
声明一个慢指针(tempIndex)记录满足>=目标数(s)的最小index,再通过不断缩小慢指针(tempIndex)与快指针(i)的长度(tempIndex++)来得出符合要求的最小长度
@method minSubArrayLen
@author pxb
@leetcode https://leetcode-cn.com/problems/minimum-size-subarray-sum
*/
func minSubArrayLen(s int, nums []int) int {
	tempIndex := 0
	tempSum := 0
	minIndex := 0
	for i := 0; i < len(nums); i++ {
		tempSum += nums[i]
		for tempSum >= s {
			if minIndex == 0 || i - tempIndex + 1 < minIndex {
				minIndex = i - tempIndex + 1
			}
			tempSum -= nums[tempIndex]
			tempIndex++
		}
	}
	return minIndex
}