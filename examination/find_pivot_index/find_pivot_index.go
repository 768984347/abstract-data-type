package find_pivot_index

/**
寻找数组的中心索引
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。
@实现原理
首先计算出整个数组的和sum(O(n)),再通过记录下循环index左侧当前的和prevSum与sum进行比对,
如果相等则当前index就是中心索引,否则sum减去当前index的值而prevSum增加当前index的值继续循环比对
需要排除数组是空,右侧的和为0的特殊情况
@author pxb
@method pivotIndex
@leetcode https://leetcode-cn.com/problems/find-pivot-index
*/
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var sum, prevSum int
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum == 0 {
		return 0
	}

	prevSum = nums[0]
	for i := 1; i < len(nums); i++ {
		if sum - nums[i] == prevSum {
			return i
		} else {
			prevSum += nums[i]
			sum -= nums[i]
		}
	}
	return -1
}