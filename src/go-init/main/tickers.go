package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at ", t.Format("2006-01-02 15:04:05"))
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
