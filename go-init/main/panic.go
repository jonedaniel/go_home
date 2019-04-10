package main

import (
	"fmt"
	"os"
)

func createdFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
func main() {
	f := createdFile("defer.txt")
	//我们使用 defer通过 closeFile 来关闭这个文件。这会在封闭函数（main）结束时执行，就是 writeFile 结束后。
	defer closeFile(f)
	writeFile(f)
	// 不像在有些语言中使用异常处理错误，在 Go 中则习惯通过返回值来标示错误
	_, err := os.Create("/tem/file")
	if err != nil {
		panic("can`t find the path.")
	}

}
