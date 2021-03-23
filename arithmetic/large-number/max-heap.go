package large_number

type MaxHeap struct {
	size int
	heap []int
}

func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		size: size,
		heap: make([]int, 0),
	}
}

func (m *MaxHeap) BuildMaxHeap(nums []int) {
	for _, n := range nums {
		m.Add(n)
	}
}

func (m *MaxHeap) Add(n int) {
	if m.size > len(m.heap) {
		m.heap = append(m.heap, n)
		m.Update(len(m.heap) - 1)
	} else if n < m.heap[len(m.heap)-1] {

	}
}

func (m *MaxHeap) Update(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if index+1 < len(m.heap) && m.heap[index] < m.heap[index+1] {
			index++
		}
		if m.heap[parent] < m.heap[index] {
			m.heap[parent], m.heap[index] = m.heap[index], m.heap[parent]
			index = parent
		} else {
			break
		}
	}
}

func (m *MaxHeap) ComeUp(index int) {
	for 2*index+1 < len(m.heap) {
		child := 2*index + 1
		if child+1 < len(m.heap) && m.heap[child+1] < m.heap[child] {
			child++
		}
		if m.heap[child] > m.heap[index] {
			m.heap[child], m.heap[index] = m.heap[index], m.heap[child]
			index = child
		} else {
			break
		}
	}
}

func (m *MaxHeap) Heap() []int {
	return m.heap
}
