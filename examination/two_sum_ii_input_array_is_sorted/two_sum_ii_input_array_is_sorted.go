package two_sum_ii_input_array_is_sorted

/**
@题目(两数之和 II - 输入有序数组)
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

@实现原理
利用map的key存储目标数和当前遍历的元素相减的number,在循环遍历中只要存在与map的key相等的数即为最后的答案
@method twoSum
@author pxb
@leetcode https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
*/
func twoSum(numbers []int, target int) []int {
	numberMap := map[int]int{}
	for i := 0; i < len(numbers); i++ {
		_, exists := numberMap[numbers[i]]
		if exists {
			return []int{numberMap[numbers[i]] + 1, i + 1}
		} else {
			numberMap[target - numbers[i]] = i
		}
	}
	return []int{}
}