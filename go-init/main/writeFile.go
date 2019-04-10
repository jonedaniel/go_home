package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	d1 := []byte("hello\ngo\n")
	_ = ioutil.WriteFile("fileToWrite", d1, 0644)

	f, _ := os.Create("ftw2")

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, _ := f.Write(d2)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, _ := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	_ = f.Sync()

	w := bufio.NewWriter(f)
	n4, _ := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n", n4)
	_ = w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
