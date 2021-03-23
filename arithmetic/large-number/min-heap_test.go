package large_number

import (
	"testing"
)

func TestMinHeap_BuildMinHeap(t *testing.T) {
	minHeap := NewMinHeap(6)
	tmp := []int{3, 5, 2, 8, 11, 4}
	minHeap.BuildMinHeap(tmp)
	minHeap.Add(4)
	t.Log(minHeap.nums, minHeap.size)
}

func TestMinHeap_BuildMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap(6)
	tmp := []int{3, 5, 2, 8, 11, 4}
	maxHeap.BuildMaxHeap(tmp)
	//minHeap.Add(4)
	t.Log(maxHeap.heap)
}
