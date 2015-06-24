// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"

	"sync"
)

func main() {
	race()
	correct()
	alsoCorrect()
}

// 이 함수는 데이터 경합을 가지고서 아마 55555를 출력하거나 다른 것을 출력할 것입니다.
// This function has a data race and may print "55555", or something else.
func race() {
	fmt.Println("Data race:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			//i가 6개의 고루틴에 공유된다.
			fmt.Print(i) // The variable i is shared by six (6) goroutines.
			wg.Done()
		}()
	}
	//5개의 고루틴이 종료되기를 기다림
	wg.Wait() // Wait for all five goroutines to finish.
	fmt.Println()
}

func correct() {
	fmt.Println("Correct:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		//지역변수 사용
		go func(n int) { // Use a local variable.
			fmt.Print(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
}

func alsoCorrect() {
	fmt.Println("Also correct:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		n := i // Create a unique variable for each closure.
		go func() {
			fmt.Print(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}
