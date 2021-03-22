package main

type MinHeap struct {
	size int
	nums []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		nums: make([]int, 0),
	}
}

func (m *MinHeap) Parent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}

func (m *MinHeap) LeftChild(i int) int {
	return i*2 + 1
}

func (m *MinHeap) RightChild(i int) int {
	return 2 * (i + 1)
}

func (m *MinHeap) BuildHeap(n int) {
	for i := 0; i < n; i++ {

	}
}
