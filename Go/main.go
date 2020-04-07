package main

import (
	"fmt"
	"gotype"
	"string"
)

// SliceMode SliceMode
func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Slice 构建栈结构")
	slicestack := gotype.NewSliceStack(make([]int, 0))
	slicestack.Push(1)
	fmt.Println("栈顶元素为：", slicestack.Top())
	fmt.Println("栈大小为：", slicestack.Size())
	slicestack.PoP()
	fmt.Println("弹栈成功：", slicestack.Size())
	slicestack.PoP()
}

// LinkedMode LinkedMode
func LinkedMode() {
	defer func () {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Linked 构建队列结构")
	linkedQueue := gotype.NewSliceQueue(make([]int, 0))
	linkedQueue.AddQueue(1)
	linkedQueue.AddQueue(2)
	fmt.Println("队列头元素为：", linkedQueue.GetFront())
	fmt.Println("队列尾元素为：", linkedQueue.GetBack())
	fmt.Println("队列大小为：", linkedQueue.Size())
}

func main() {
	// SliceMode()
	// LinkedMode()
	string.PermutationStr("aaaa")
}
