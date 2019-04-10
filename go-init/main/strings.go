package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	p("Contains	Count	HasPrefix	HasSuffix	Index	Join" +
		"	Repeat	Split	ToLower	ToUpper")
	//全部替换
	p("Replace", strings.Replace("foo", "o", "0", -1))
	//只替换第一个
	p("Replace", strings.Replace("foo", "o", "0", 1))

	p("Len:", len("hello"))
	p("Char:", "hello"[1])
	//ASCII to char
	p(string("hello"[1]))
	//char to ASCII
	p(int('e'))
}
