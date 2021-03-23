package large_number

type MinHeap struct {
	size int
	nums []int
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		nums: make([]int, 0),
		size: size,
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

func (m *MinHeap) BuildMinHeap(nums []int) {
	for _, n := range nums {
		m.Add(n)
	}
}

func (m *MinHeap) Add(n int) {
	if m.size > len(m.nums) {
		m.nums = append(m.nums, n)
		m.Update(len(m.nums) - 1)
	} else if n > m.nums[0] {
		m.nums[0] = n
		m.Down(0)
	}
}

func (m *MinHeap) Update(index int) {
	for index > 0 {
		parent := (index - 1) >> 2
		if m.nums[parent] > m.nums[index] {
			m.nums[parent], m.nums[index] = m.nums[index], m.nums[parent]
			index = parent
		} else {
			break
		}
	}
}

func (m *MinHeap) Down(index int) {
	for 2*index+1 < len(m.nums) {
		child := 2*index + 1
		if child+1 < len(m.nums) && m.nums[child+1] < m.nums[child] {
			child++
		}
		if m.nums[index] > m.nums[child] {
			m.nums[child], m.nums[index] = m.nums[index], m.nums[child]
			index = child
		} else {
			break
		}
	}
}
