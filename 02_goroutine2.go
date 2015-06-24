// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second) // wait 1 sec for other goroutine to finish
}
