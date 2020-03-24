package example

import (
	"fmt"
	"heapq-go/heapq"
)

// Here is a example for how to code a custom struct heap.
// First,you need to implement all functions of the Interface.

// A simple []int heap
type heap struct {
	list []int
}
// Implement the interface
func (h heap) At(i int) interface{} { return h.list[i] }
// To change the []int, you need to use a struct point
func (h *heap) Assign(i int,item interface{}) {
	// interface{} needs type assertion
	if v,ok := item.(int);ok{
		h.list[i] = v
	}
}
func (h heap) Len() int { return len(h.list) }
func (h *heap) Append(item interface{}) {
	if v,ok := item.(int);ok {
		h.list = append(h.list,v)
	}
}
func (h heap) Less(i,j int) bool {return h.list[i] < h.list[j] }
func (h *heap) Pop(i int) interface{} {
	item := h.list[i]
	h.list = append(h.list[:i],h.list[i+1:]...)
	return item
}
func (h *heap) Swap(i,j int) { h.list[i],h.list[j] = h.list[j],h.list[i]}
func (h heap)  IsMinHeap() bool { return true }

// Now, you can use your heap like this.
func main(){
	// Create your heap, the best choice is use point
	myHeap := &heap{list:make([]int,0,10)}

	// Push some elements ...
	heapq.HeapPush(myHeap,3)
	heapq.HeapPush(myHeap,2)
	heapq.HeapPush(myHeap,4)
	heapq.HeapPush(myHeap,7)

	// Pop the smallest element
	smallest := heapq.HeapPop(myHeap)
	fmt.Println(smallest)

	// Or you can user Heapify() to transform a unordered heap
	unorderedHeap := &heap{list: []int{6,2,1,3,5,4}}
	heapq.Heapify(unorderedHeap)

	//...
}
