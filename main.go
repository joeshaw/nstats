package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type minHeap []float64

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Root() float64      { return h[0] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type maxHeap struct {
	minHeap
}

func (h maxHeap) Less(i, j int) bool { return h.minHeap.Less(j, i) }

func main() {
	var minh minHeap
	var maxh maxHeap

	heap.Init(&minh)
	heap.Init(&maxh)

	var n int
	var min, max, sum, sqSum float64

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		f, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, `Skipping invalid input "%s"`, s.Text())
			continue
		}

		if n == 0 || f < min {
			min = f
		}

		if n == 0 || f > max {
			max = f
		}

		n++
		sum += f
		sqSum += f * f

		if maxh.Len() == 0 || f < maxh.Root() {
			heap.Push(&maxh, f)
		} else {
			heap.Push(&minh, f)
		}

		if maxh.Len()-minh.Len() > 1 {
			x := heap.Pop(&maxh)
			heap.Push(&minh, x)
		} else if minh.Len()-maxh.Len() > 1 {
			x := heap.Pop(&minh)
			heap.Push(&maxh, x)
		}
	}

	fmt.Printf("N       %d\n", n)

	if n == 0 {
		return
	}

	fmt.Printf("min     %g\n", min)
	fmt.Printf("max     %g\n", max)
	fmt.Printf("sum     %g\n", sum)

	var median float64
	if maxh.Len() > minh.Len() {
		median = maxh.Root()
	} else if minh.Len() > maxh.Len() {
		median = minh.Root()
	} else {
		median = (maxh.Root() + minh.Root()) / 2
	}

	fmt.Printf("median  %g\n", median)

	mean := sum / float64(n)
	fmt.Printf("mean    %g\n", mean)

	if n > 1 {
		stddev := math.Sqrt((float64(n)*sqSum - sum*sum) / float64(n*(n-1)))
		fmt.Printf("stddev  %g\n", stddev)
	}
}
