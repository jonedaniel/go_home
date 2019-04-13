package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println

	p(time.Now())
	p(time.Now().Location())
	then := time.Date(2019, 04, 9, 20, 34, 58, 651387237, time.Local)
	p(then)
	p(then.Year())
	p(then.Weekday())
	p(then.Location())

	diff := time.Now().Sub(then)
	p(diff.Hours())
	p(then.Add(diff))

	//默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，需要给定一个变化的种子。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",", rand.Intn(100))

}
