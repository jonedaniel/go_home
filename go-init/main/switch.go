package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	}
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Monday:
		println("星期一")
	case time.Tuesday:
		println("星期二")
	case time.Thursday:
		println("星期三")
	case time.Wednesday:
		println("星期四")
	case time.Friday:
		println("星期五")
	case time.Saturday, time.Sunday:
		println("周末")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println(t, " is a bool")
		case int:
			println(t, " is a int")
		case string:
			println(t, " is a string")
		default:
			println("no one know, no one cares", t)
		}
	}
	whatAmI(true)
	whatAmI("hey")
	whatAmI(1)
}
