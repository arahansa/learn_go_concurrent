// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
)

type Sushi string

func main() {
	var ch <-chan Sushi = Producer()
	for s := range ch {
		fmt.Println("먹은 것 : ", s)
	}
}

func Producer() <-chan Sushi {
	ch := make(chan Sushi)
	go func() {
		ch <- Sushi("광어회") // Ebi nigiri
		ch <- Sushi("우럭회") // Toro nigiri
		close(ch)
	}()
	return ch
}
