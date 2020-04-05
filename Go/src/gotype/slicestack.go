package gotype

import (
	"errors"
)

// SliceStack 栈
type SliceStack struct {
	arr       []int
	stackSize int
}

// IsEmpty 是否为空
func (p *SliceStack) IsEmpty() bool {
	return p.stackSize == 0
}

// Size 栈大小
func (p *SliceStack) Size() int {
	return p.stackSize
}

// Top 栈顶
func (p *SliceStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空"))
	}
	return p.arr[p.stackSize-1]
}

// PoP 弹出
func (p *SliceStack) PoP() int {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	panic(errors.New("栈已经为空"))
}

// Push 推进栈
func (p *SliceStack) Push(t int) {
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}

// NewSliceStack NewSliceStack
func NewSliceStack(arr []int) *SliceStack{
	// slicestack := new(SliceStack)
	// slicestack.arr = arr;
	slicestack := &SliceStack{ arr: arr }
	return slicestack
}
