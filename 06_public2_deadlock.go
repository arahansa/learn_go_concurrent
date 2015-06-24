// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

func main() {
	wait := Publish("채널은 고루틴들이 통신하게 해줍니다..", 2*time.Second)
	fmt.Println("뉴스를 기다려볼까요")
	<-wait
	fmt.Println("뉴스가 끝났습니다. 갈 시간이네요!.")
}

// Publish는 주어진 시간 후에 텍스트를  stdout에 출력합니다.
// 텍스트가 발행된 후, wait채널을 닫습니다.

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("뉴스 속보:", text)
		//close(ch) // 닫힌 채널은 영원히 제로값을 전송합니다.
	}()
	return ch
}
