// http://www.nada.kth.se/~snilsson/concurrency/
package main

import "fmt"

func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// At this point the program execution stops and all
	// active goroutines are killed.
}
