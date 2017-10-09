package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	var aa chan<- int = c // send-only
	var bb <-chan int = c // receive-only
	aa <- 1
	//<-send // Error: receive from send-only type chan<- int
	<-bb
	// recv <- 2 // Error: send to receive-only type <-chan int
	fmt.Printf("finished!!!\n")
	var tm = time.Now()
	fmt.Printf("now:%v\n",tm)
}
