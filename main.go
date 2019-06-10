package main

import (
	stack2 "abstract_data_type/stack_LIFO"
	"fmt"
)

func main() {
	stack := stack2.Init(3)
	suc := stack.Push(2)
	fmt.Println("2 push status = ", suc)
	num, err := stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)
	suc = stack.Push(3)
	fmt.Println("3 push status = ", suc)
	suc = stack.Push(4)
	fmt.Println("4 push status = ", suc)
	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)

	suc = stack.Push(5)
	fmt.Println("5 push status = ", suc)
	suc = stack.Push(6)
	fmt.Println("6 push status = ", suc)
	suc = stack.Push(7)
	fmt.Println("7 push status = ", suc)

	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)

	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)

	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)

	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)

	num, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop num is:", num)
}