package pq

// Heap definitions
// The binary heap is a data structure that can efficiently support the basic
// priority-queue operations. In a binary heap, the items are sorted in an array
// such that each key is guaranteed to be larger than (or equal to) the keys at
// two other specific positions. In turn, each of those keys must be larger than
// two more keys, and so forth. This ordering is easy to see if we view the keys
// as being in a binary tree structure with edges from each key to the two keys
// known to be smaller.

// Definition: A binary tree is heap-ordered if the key in each node is larger
// than (or equal to) the keys in that nodes two children (if any).

// Proposition
// The largest key in a heap-ordered binary tree is found at the root.

// Definition
// A binary heap is a set of nodes with keys arranged in a complete heap-ordered
// binary tree, represented in level order in an array (not using the first position).

// Algorithms on heaps.
// We represent a heap of size n in private array arr[] of length n+1, with arr[0]
// unused and the heap in arr[1] through arr[n]. We access keys only through private
// helper functions Less() and Swap(). The heap operations that we consider work by
// first making a simple modification that could violate the heap condition, then
// traveling through the heap, modifying the heap as required to ensure that the
// heap condition is satisfied everywhere. We refer to this process as heapify.

// Bottom-up heapify (swim).
// If the heap order is violated because a node's key becomes larger than that
// node's parent key, then we can make progress toward fixing the violation by
// exchanging the node with its parent. After the exchange, the node is larger
// than both its children (one is the old parent, and the other is smaller than
// the old parent because it was a child of that node) but the node may still be
// larger than its parent. We can fix that violation in the same way, and so forth,
// moving up the heap until we reach a node with a larger key, or the root.

// Top-down heapify (sink).
// If the heap order is violated because a node's key becomes smaller than one
// or both of that node's children's keys, then we can make progress toward fixing
// the violation by exchanging the node with the larger of its two children. This
// switch may cause a violation at the child. We fix that violation in the same
// way, and so forth, moving down the heap until we reach a node with both child
// smaller, or bottom.

type BinaryHeap interface {
	GetLeftChildIndex(p, n int) int
	GetRightChildIndex(p, n int) int
	GetParentIndex(c, n int) int
	GetRootIndex() int
	GetLastLeafIndex(n int) int
}

// default 1-based indexing
func NewBinaryHeap() BinaryHeap {
	return BinaryHeapBased1{}
}

func swap(a []Item, i, j int) {
	a[i], a[j] = a[j], a[i]
}
