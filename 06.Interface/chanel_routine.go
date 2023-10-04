package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	c <- payload
}

func main() {
	ch := make(chan int)
	go routine1(ch, 1)

	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
