package pascals_triangle

/**
@题目(118. 杨辉三角)
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

@实现原理
将给定数设置为for循环次数,获取当前index左上方和右上方的数相加,如果取不到或者超过最大index忽略即可
@method generate
@author pxb
@leetcode https://leetcode-cn.com/problems/pascals-triangle/
*/
func generate(numRows int) [][]int {
	var returnArr [][]int
	for i := 0; i < numRows; i++ {
		//初始化行
		returnArr = append(returnArr, []int{})
		for j := 0; j <= i; j++ {
			leftIndex := j - 1
			rightIndex := j
			leftNum := 0
			rightNum := 0
			if i > 0 {
				if leftIndex >= 0 {
					leftNum = returnArr[i - 1][leftIndex]
				}
				if rightIndex < i {
					rightNum = returnArr[i - 1][rightIndex]
				}
			} else {
				leftNum = 1
			}
			returnArr[i] = append(returnArr[i], leftNum + rightNum)
		}
	}
	return returnArr
}