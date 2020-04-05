package main

import (
	"fmt"
	"gotype"
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
	fmt.Println(slicestack)
}

func main() {
	SliceMode()
}
