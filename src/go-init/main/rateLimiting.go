package main

import (
	"fmt"
	"time"
)

func main() {
	reqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		reqs <- i
	}
	close(reqs)
	limiter := time.Tick(time.Second * 2)

	for r := range reqs {
		<-limiter
		fmt.Println("req", r, " time:", time.Now())
	}

	//脉冲限制
	burstyReqs := make(chan time.Time, 5)
	burstyReqs <- time.Now()
	burstyReqs <- time.Now()
	burstyReqs <- time.Now()

	go func() {
		for t := range time.Tick(time.Second * 2) {
			burstyReqs <- t
		}
	}()

	for i := 0; i < 5; i++ {
		<-burstyReqs
		fmt.Println("time:", time.Now())
	}

}
