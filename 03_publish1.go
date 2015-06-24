// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

func main() {
	Publish("고루틴이 새로운 쓰레드실행을 시작합니다..", 2*time.Second)
	fmt.Println("제가 끝내기 전에 새로운 뉴스가 오길 기다려볼까요.")

	// Wait for the news to be published.
	time.Sleep(5 * time.Second)

	fmt.Println("5초 후 : 끝낼게요.")
}
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("새로운 뉴스:", text)
	}() // Note the parentheses. We must call the anonymous function.
}
