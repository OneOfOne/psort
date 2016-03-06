package psort

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type pQS struct {
	sync.WaitGroup
	Interface
	numWorkers int32
}

func (p *pQS) sort(a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(p, a, b)
			p.Done()
			return
		}
		maxDepth--
		mlo, mhi := doPivot(p, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		bg := atomic.AddInt32(&p.numWorkers, -1) > -1
		if mlo-a < b-mhi {
			if bg {
				p.Add(1)
				go p.sort(a, mlo, maxDepth)
			} else {
				quickSort(p, a, mlo, maxDepth)
			}
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			if bg {
				p.Add(1)
				go p.sort(mhi, b, maxDepth)
			} else {
				quickSort(p, mhi, b, maxDepth)
			}
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if p.Less(i, i-6) {
				p.Swap(i, i-6)
			}
		}
		insertionSort(p, a, b)
	}
	p.Done()
}

func pSort(data Interface, n, maxDepth int) {
	p := &pQS{
		Interface:  data,
		numWorkers: int32(runtime.NumCPU() * 2),
	}
	p.Add(1)
	p.sort(0, n, maxDepth)
	p.Wait()
}
