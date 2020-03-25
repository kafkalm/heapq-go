package heapq

// This file provides some base-heaps.

// a []int heap, you can set `isMinHeap` to choose it is min or max.
type IntHeap struct {
	heap []int
	isMinHeap bool
}
func (h *IntHeap) Append(item interface{}) {
	if v,ok := item.(int);ok {
		h.heap = append(h.heap,v)
	}
}
func (h *IntHeap) Assign(i int, item interface{}) {
	if v,ok := item.(int);ok {
		h.heap[i] = v
	}
}
func (h *IntHeap) At(i int) interface{} {return h.heap[i]}
func (h *IntHeap) IsMinHeap() bool {return h.isMinHeap}
func (h *IntHeap) Len() int {return len(h.heap)}
func (h *IntHeap) Less(i, j int) bool {return h.heap[i] < h.heap[j]}
func (h *IntHeap) Pop(i int) interface{} {
	items := h.heap[i]
	h.heap = append(h.heap[:i],h.heap[i+1:]...)
	return items
}
func (h *IntHeap) Swap(i, j int) {h.heap[i],h.heap[j] = h.heap[j],h.heap[i]}


// a []string heap, you can set `isMinHeap` to choose it is min or max.
type StringHeap struct {
	heap []string
	isMinHeap bool
}
func (h *StringHeap) Append(item interface{}) {
	if v,ok := item.(string);ok {
		h.heap = append(h.heap,v)
	}
}
func (h *StringHeap) Assign(i int, item interface{}) {
	if v,ok := item.(string);ok {
		h.heap[i] = v
	}
}
func (h *StringHeap) At(i int) interface{} {return h.heap[i]}
func (h *StringHeap) IsMinHeap() bool {return h.isMinHeap}
func (h *StringHeap) Len() int {return len(h.heap)}
func (h *StringHeap) Less(i, j int) bool {return h.heap[i] < h.heap[j]}
func (h *StringHeap) Pop(i int) interface{} {
	items := h.heap[i]
	h.heap = append(h.heap[:i],h.heap[i+1:]...)
	return items
}
func (h *StringHeap) Swap(i, j int) {h.heap[i],h.heap[j] = h.heap[j],h.heap[i]}
