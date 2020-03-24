# Introducing

`heapq-go` is a golang version for the python library `heapq`.  
It has implemented a series of functions which are provided by `heapq`, such as `heappush()`, `heappop()`, `heapify()` and so on.  

# Installation

```
$ go get -u github.com/kafkalm/heapq-go
```

# Features

You can create your own heap by implementing the `Interface` which is defined in `heapq-go/heapq-go.go`.  
When you implement the interface, you can use all functions provided in `heapq-go/heapq-go.go`.  

# Usage

The directory `example` has a example, which implements a []int min heap.