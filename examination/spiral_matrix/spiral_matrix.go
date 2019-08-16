package spiral_matrix

/**
@题目(54. 螺旋矩阵)
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

@实现原理
顺时针螺旋循环数组遇到边界情况对方向和最大最小row,col进行修正即可
@method spiralOrder
@author pxb
@leetcode https://leetcode-cn.com/problems/spiral-matrix
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	return_arr := []int{}
	minRow := 0
	maxRow := len(matrix) - 1
	currRow := 0
	minCol := 0
	maxCol := len(matrix[0]) - 1
	currCol := 0
	rowIncr := 0
	colIncr := 1
	length := (maxRow + 1) * (maxCol + 1)
	for i := 0; i < length; i++ {
		return_arr = append(return_arr, matrix[currRow][currCol])
		if rowIncr == 0 && colIncr > 0 {
			//向右
			if colIncr + currCol > maxCol {
				//到达边缘向下修正
				//将方向变为向下
				rowIncr = 1
				colIncr = 0
				minRow++
			}
		} else if rowIncr == 0 && colIncr < 0 {
			//向左
			if colIncr + currCol < minCol {
				//到达边缘向上修正
				//将方向变为向上
				rowIncr = -1
				colIncr = 0
				maxRow--
			}
		} else if rowIncr > 0 && colIncr == 0 {
			//向下
			if rowIncr + currRow > maxRow {
				//到达边缘向左修正
				//将方向变为向左
				rowIncr = 0
				colIncr = -1
				maxCol--
			}
		} else {
			//向上
			if rowIncr + currRow < minRow {
				//到达边缘向右修正
				//将方向变为向右
				rowIncr = 0
				colIncr = 1
				minCol++
			}
		}
		currRow += rowIncr
		currCol += colIncr
	}
	return return_arr
}