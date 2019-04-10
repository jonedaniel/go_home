package main

import (
	"fmt"
	"sort"
)

type ByLength []string

//此处实现了sort.Interface的方法
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	xxx := []string{"gotohell", "fuckyo", "screwyou"}
	sort.Sort(ByLength(xxx))
	fmt.Println(xxx)

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("sorted ints:	", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("is sorted? :", s)
}
