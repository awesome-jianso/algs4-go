package heap

type MinHeap struct {
	maxN  int // maximum number of elements
	n     int // number of elements
	items []Comparable
	BinaryHeap
}

// 1-based indexing VS 0-based: almost the same performance
// see bench_test.go
func NewMinHeap() *MinHeap {
	items := make([]Comparable, defaultMaxSize)

	return &MinHeap{defaultMaxSize, 0, items, NewBinaryHeapBased1()}
}

// only for test
func NewMinHeapBased0() *MinHeap {
	items := make([]Comparable, defaultMaxSize)

	return &MinHeap{defaultMaxSize, 0, items, NewBinaryHeapBased0()}
}

func NewMinHeapBased1() *MinHeap {
	items := make([]Comparable, defaultMaxSize+1)

	return &MinHeap{defaultMaxSize, 0, items, NewBinaryHeapBased1()}
}

func (h *MinHeap) Insert(item Comparable) {
	if h.IsFull() {
		panic("heap is full")
	}

	lastLeaf := h.GetLastLeafIndex(h.n)

	h.items[lastLeaf+1] = item
	h.Swim(lastLeaf+1, lastLeaf+1)

	h.n++
}

func (h *MinHeap) Remove() Comparable {
	if h.IsEmpty() {
		panic("heap is empty")
	}

	item := h.GetHighestPriority()
	root := h.GetRootIndex()
	lastLeaf := h.GetLastLeafIndex(h.n)

	h.swap(root, lastLeaf)
	h.Sink(root, lastLeaf-1)

	h.n--

	return item
}

func (h *MinHeap) GetHighestPriority() Comparable {
	return h.items[0]
}

func (h *MinHeap) Sink(p, max int) {

	smallerChild := h.getSmallerChildIndex(p, max)

	// if the left and right child do not exist
	// stop sinking
	if smallerChild == -1 {
		return
	}

	// compare the smaller child with the current index to see
	// if swap and further sink is needed
	if h.less(smallerChild, p) {
		h.swap(smallerChild, p)
		h.Sink(smallerChild, max)
	}

}

func (h *MinHeap) Swim(c, max int) {
	parent := h.GetParentIndex(c, max)

	if parent != -1 && h.less(c, parent) {
		h.swap(c, parent)
		h.Swim(parent, max)
	}
}

func (h *MinHeap) IsFull() bool {
	return h.n == h.maxN
}

func (h *MinHeap) IsEmpty() bool {
	return h.n == 0
}

func (h *MinHeap) GetItems() []Comparable {
	begin := h.GetRootIndex()
	end := h.GetLastLeafIndex(h.n)
	return h.items[begin : end+1]
}

// find the minimum of the left and right child elements
// if do not exist, return -1
func (h *MinHeap) getSmallerChildIndex(p, max int) int {
	leftChild := h.GetLeftChildIndex(p, max)
	rightChild := h.GetRightChildIndex(p, max)

	if leftChild != -1 && rightChild != -1 {
		if h.less(leftChild, rightChild) {
			return leftChild
		} else {
			return rightChild
		}
	} else if leftChild != -1 {
		return leftChild
	} else if rightChild != -1 {
		return rightChild
	} else {
		return -1
	}
}

func (h *MinHeap) getItemAtIndex(i int) Comparable {
	return h.items[i]
}

// if items[i] < items[j]
func (h *MinHeap) less(i, j int) bool {
	one := h.getItemAtIndex(i)
	another := h.getItemAtIndex(j)
	return one.CompareTo(another) < 0
}

func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
