// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	race()
	sharingIsCaring()
	lockItUp()
}

// 이 함수는 데이터레이스를 가지고 있고, 아마 1을 출력하거나 다른 것을 출력할 것입니다.
// This function has a data race and may print "1", or something else.
func race() {
	fmt.Println("Wrong (data race):")
	wg := new(sync.WaitGroup)
	a := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			a = a + 1 // one access: read, increment, write
			wg.Done()
		}()
		a = a - 1
	}
	wg.Wait()
	fmt.Println(a) // Output: UNSPECIFIED
}

// This is the preferred way to handle concurrent data access in Go:
// "Don't communicate by sharing memory; share memory by communicating."
func sharingIsCaring() {
	fmt.Println("Good (share memory by communicating):")
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another goroutine.
	n++
	fmt.Println(n) // Output: 2
}

// Sometimes it's more appropriate to use explicit locking.
// This can often be handled safely and transparently by building
// a data structure that does the synchronization internally.
func lockItUp() {
	fmt.Println("Good (explicit locking):")
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // one access
		close(wait)
	}()
	n.Add(1) // another concurrent access
	<-wait
	fmt.Println(n.Value()) // Output: 2
}

// AtomicInt is a concurrent datastructure that holds an int.
// Its zero value is 0.
type AtomicInt struct {
	mu sync.Mutex // A lock than can be held by just one goroutine at a time.
	n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
	a.mu.Lock() // Wait for the lock to be free and then take it.
	a.n += n
	a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}
