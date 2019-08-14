package diagonal_traverse

/**
对角线遍历
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素
@实现原理
当对角线方向为右上时如果col超过边界需要向下修正,如果row少于0行需要向右修正,其余情况属于正常的找下一个对角点,对角线方向为左下时与右上相反
@author pxb
@method findDiagonalOrder
@method findNext
@leetcode https://leetcode-cn.com/problems/diagonal-traverse/
*/
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var returnArr []int
	lastRowIndex := 0
	lastColIndex := 0
	incr := -1
	count := len(matrix) * len(matrix[0])
	for i := 0; i < count; i++ {
		returnArr = append(returnArr, matrix[lastRowIndex][lastColIndex])
		lastColIndex, lastRowIndex, incr = findNext(len(matrix), len(matrix[0]), lastRowIndex, lastColIndex, incr)
	}
	return returnArr
}

func findNext(maxRowIndex int, maxColIndex int, currRowIndex int, currColIndex int, incr int) (nextCol int, nextRow int, currIncr int) {
	if incr > 0 {
		//对角线方向为左下↙
		if currRowIndex + 1 >= maxRowIndex {
			nextRow = currRowIndex
			nextCol = currColIndex + 1
			//到达边界改变对角线方向
			incr *= -1
		} else if currColIndex - 1 < 0 {
			nextRow = currRowIndex + 1
			nextCol = currColIndex
			//到达边界改变对角线方向
			incr *= -1
		} else {
			nextRow = currRowIndex + 1
			nextCol = currColIndex - 1
		}
	} else {
		//对角线方向为右上↗
		if currColIndex + 1 >= maxColIndex {
			nextRow = currRowIndex + 1
			nextCol = currColIndex
			//到达边界改变对角线方向
			incr *= -1
		} else if currRowIndex - 1 < 0 {
			nextRow = currRowIndex
			nextCol = currColIndex + 1
			//到达边界改变对角线方向
			incr *= -1
		} else {
			nextRow = currRowIndex - 1
			nextCol = currColIndex + 1
		}
	}
	return nextCol, nextRow, incr
}