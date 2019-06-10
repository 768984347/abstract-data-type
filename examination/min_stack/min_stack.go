package min_stack

import "fmt"

/**
最小栈
@特性
先进后出 支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
@实现原理
只需要保留当前数组的len作为下一个插入的元素index即可
由于不限制栈的大小当超出切片长度时需要使用append来扩展切片
@method Constructor 初始化栈
@method Push 入栈
@method Pop 出栈
@method IsEmpty 判断栈是否为空
@method GetMin O(n)获取最小元素
@method GetLength 获取栈长度
@method Top 获取栈首个元素
@leetcode https://leetcode-cn.com/explore/learn/card/queue-stack/218/stack-last-in-first-out-data-structure/877/
*/

type MinStack struct {
	data []int
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	//如果已经超出了切片的长度需要append来增加长度
	if this.length >= len(this.data) {
		this.data = append(this.data, x)
	} else {
		this.data[this.length] = x
	}
	this.length++
}

func (this *MinStack) Pop() {
	if !this.IsEmpty() {
		this.length--;
	}
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return -1
	}
	fmt.Println(this.data)
	fmt.Println(this.length)
	return this.data[this.length - 1]
}

func (this *MinStack) GetMin() int {
	//先取出最上面的元素默认是最小的
	min_data := this.Top()
	for i := 0; i < this.length - 1; i++ {
		if this.data[i] < min_data {
			min_data = this.data[i]
		}
	}
	return min_data
}

func (this *MinStack) IsEmpty() bool {
	return this.length <= 0
}