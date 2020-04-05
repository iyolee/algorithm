package gotype

import (
	"errors"
)

// SliceQueue 队列
type SliceQueue struct {
	arr[] int
	front int
	rear int
}

// IsEmpty IsEmpty
func (p *SliceQueue) IsEmpty() bool {
	return p.front == p.rear
}

// Size Size
func (p *SliceQueue) Size() int {
	return p.rear - p.front
}

// GetFront 队列首元素
func (p *SliceQueue) GetFront() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	return p.arr[p.front]
}

// GetBack 队列尾元素
func (p *SliceQueue) GetBack() int {
	if (p.IsEmpty()) {
		panic(errors.New("队列已经为空"))
	}
	return p.arr[p.rear - 1]
}

// DelQueue 删除队头元素
func (p *SliceQueue) DelQueue() {
	if (p.rear > p.front) {
		p.rear--
		p.arr = p.arr[1:]
	} else {
		panic(errors.New("队列已经为空"))
	}
}

// AddQueue 把新元素加入队尾
func (p *SliceQueue) AddQueue(item int) {
	p.arr = append(p.arr, item)
	p.rear++
}

// NewSliceQueue NewSliceQueue
func NewSliceQueue(arr []int) *SliceQueue{
	// slicestack := new(SliceStack)
	// slicestack.arr = arr;
	slicequeue := &SliceQueue{ arr: arr }
	return slicequeue
}