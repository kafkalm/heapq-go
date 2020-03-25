# Introducing

`heapq-go` is a golang version of the python library `heapq`.  
It has implemented a series of functions which are provided by `heapq`, such as `heappush()`, `heappop()`, `heapify()` and so on.  

# Installation

``` shell
$ go get -u github.com/kafkalm/heapq-go
```

# Features

You can create your own heap by implementing the `Interface` which is defined in `heapq/heapq.go`.  
When you implement the interface, you can use all functions provided in `heapq/heapq.go`.  

# Usage

The directory `example` has a example, which implements a []int min heap.  
The file `heapq/heaps.go` provides some base-heaps, like []int heap, []string heap, and in future, it will provides more basic type heaps.  
When you want to use the heaps in `heapq/heaps.go`, it's easy to create a heap,like this

``` go
import . "github.com/kafkalm/heapq-go/heapq"

func main(){
    // if MinFlag == true, means the heap is a min heap.
    myheap := &IntHeap{Heap:make([]int,0,10),MinFlag:true}

    HeapPush(myheap,1)
    HeapPush(myheap,2)

    value := HeapPop(myheap)
    if v,ok := value.(int);ok{
        do_something(v)
        ...
    }

    unorderedHeap := &StringHeap{Heap:[]string{"Hello","Heapq"},MinFlag:true}

    // Make the unordered heap to a ordered heap
    Heapify(unorderedHeap)

}
```