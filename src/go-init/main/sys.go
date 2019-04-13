package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	//fmt.Println(arg)

	for _, e := range os.Environ() {
		pair := strings.Split(e, "-")
		fmt.Println(pair[0])
	}
}
