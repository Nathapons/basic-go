package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Run1 Something")
	}
}

func run2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Run2 Something")
	}
}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
