package heapq

// Heapq : A python library that provides heap sort algorithm and some functions about heap/heap queue.
// Implemented Golang version by kafkal


type Interface interface {
	Append(item interface{})	// append a item into your heap
	Assign(i int,item interface{})	// assign the i'th as item, just like slice[i] = item
	At(i int) interface{}		// return the i'th item, just like slice[i]
	IsMinHeap() bool			// if your heap need to be a min heap, return true,
								// otherwise, return false
	Len() int					// return your heap's length
	Less(i,j int) bool			// compare your heap's i'th item and j'th item.
								// if item_i < item_j, return true , else return false
	Pop(i int) interface{}		// pop the i'th item from your heap
	Swap(i,j int)				// swap the i'th item and the j'th item
}


// Here are all functions you can use in your heap
// when you have already implemented all interface
// fucntions.


// Push item into your heap
func HeapPush(heap Interface,item interface{}){
	heap.Append(item)
	siftdown(heap,0,heap.Len()-1,heap.IsMinHeap())
}


// Pop the smallest item from your heap
func HeapPop(heap Interface) (item interface{}) {
	if heap.Len() <= 0 {
		return nil
	} else {
		tmp := heap.Pop(heap.Len()-1)
		if heap.Len() > 0{
			item = heap.At(0)
			heap.Assign(0,tmp)
			siftup(heap,0,heap.IsMinHeap())
			return item
		}
		return tmp
	}
}


// Make a initialize, unordered heap to ordered
// in-place, in O(heap.Len()) time
func Heapify(heap Interface) {
	n := heap.Len()
	for i:=n/2-1;i>=0;i-- {
		siftup(heap,i,heap.IsMinHeap())
	}
}


// Pop and return the current smallest value, and add the item.
// This function is more efficient than HeapPop() followed by HeapPush()
func HeapReplace(heap Interface,item interface{}) (returnitem interface{}) {
	if heap.Len() <= 0 {
		return nil
	} else {
		returnitem = heap.At(0)
		heap.Assign(0,item)
		siftup(heap,0,heap.IsMinHeap())
		return returnitem
	}
}


func siftdown(heap Interface,startpos,pos int,flag bool) {
	for ;pos > startpos; {
		parentpos := (pos - 1) >> 1
		if flag == heap.Less(pos,parentpos){
			heap.Swap(pos,parentpos)
			pos = parentpos
			continue
		}
		break
	}
}


func siftup(heap Interface,pos int,flag bool) {
	endpos := heap.Len()
	startpos := pos
	newitem := heap.At(pos)
	childpos := 2*pos + 1
	for ;childpos < endpos; {
		rightpos := childpos + 1
		if rightpos < endpos && (flag != heap.Less(childpos,rightpos)) {
			childpos = rightpos
		}
		heap.Assign(pos,heap.At(childpos))
		pos = childpos
		childpos = 2 * pos + 1
	}
	heap.Assign(pos,newitem)
	siftdown(heap,startpos,pos,flag)
}
