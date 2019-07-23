package daily_temperatures

/**
每日温度
根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高的天数。如果之后都不会升高，请输入 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
@实现原理
利用栈的后进先出原理在循环中比对栈顶的大小(如果栈为空将给定切片的第一个元素入栈作为栈顶),如果大于栈顶则栈顶出栈将index差(天数)记录否则入栈继续寻找下一个大于栈顶的元素
@author pxb
@method dailyTemperatures
@leetcode https://leetcode-cn.com/problems/daily-temperatures
*/
func dailyTemperatures(T []int) []int {
	var res, temp, tempIndex []int
	for i := 0; i < len(T); i++ {
		if len(temp) == 0 || T[i] <= temp[len(temp) - 1] {
			temp = append(temp, T[i])
			tempIndex = append(tempIndex, i)
			res = append(res, 0)
			continue
		} else {
			if i - tempIndex[len(temp) - 1] >= 1 {
				res[tempIndex[len(temp) - 1]] = i - tempIndex[len(temp) - 1]
			} else {
				res = append(res, 1)
			}
			temp = temp[:len(temp) - 1]
			tempIndex = tempIndex[:len(tempIndex) - 1]
			i--
		}
	}

	return res
}