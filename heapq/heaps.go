package heapq

// This file provides some base-Heaps.

// a []int heap, you can set `MinFlag` to choose it is min or max.
// MinFlag = true, means the heap is a min heap.
type IntHeap struct {
	Heap []int
	MinFlag bool
}
func (h *IntHeap) Append(item interface{}) {
	if v,ok := item.(int);ok {
		h.Heap = append(h.Heap,v)
	}
}
func (h *IntHeap) Assign(i int, item interface{}) {
	if v,ok := item.(int);ok {
		h.Heap[i] = v
	}
}
func (h *IntHeap) At(i int) interface{} {return h.Heap[i]}
func (h *IntHeap) IsMinHeap() bool {return h.MinFlag}
func (h *IntHeap) Len() int {return len(h.Heap)}
func (h *IntHeap) Less(i, j int) bool {return h.Heap[i] < h.Heap[j]}
func (h *IntHeap) Pop(i int) interface{} {
	items := h.Heap[i]
	h.Heap = append(h.Heap[:i],h.Heap[i+1:]...)
	return items
}
func (h *IntHeap) Swap(i, j int) {h.Heap[i],h.Heap[j] = h.Heap[j],h.Heap[i]}


// a []string heap, you can set `MinFlag` to choose it is min or max.
type StringHeap struct {
	Heap []string
	MinFlag bool
}
func (h *StringHeap) Append(item interface{}) {
	if v,ok := item.(string);ok {
		h.Heap = append(h.Heap,v)
	}
}
func (h *StringHeap) Assign(i int, item interface{}) {
	if v,ok := item.(string);ok {
		h.Heap[i] = v
	}
}
func (h *StringHeap) At(i int) interface{} {return h.Heap[i]}
func (h *StringHeap) IsMinHeap() bool {return h.MinFlag}
func (h *StringHeap) Len() int {return len(h.Heap)}
func (h *StringHeap) Less(i, j int) bool {return h.Heap[i] < h.Heap[j]}
func (h *StringHeap) Pop(i int) interface{} {
	items := h.Heap[i]
	h.Heap = append(h.Heap[:i],h.Heap[i+1:]...)
	return items
}
func (h *StringHeap) Swap(i, j int) {h.Heap[i],h.Heap[j] = h.Heap[j],h.Heap[i]}
