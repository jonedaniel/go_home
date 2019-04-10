package main

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func main() {
//	dat, e := ioutil.ReadFile("defer.txt")
//	check(e)
//	fmt.Print(string(dat))
//	f, err := os.Open("defer.txt")
//	check(err)
//
//	b1 := make([]byte, 5)
//	n1,err := f.Read(b1)
//	check(err)
//	fmt.Printf("%d bytes: %s\n",n1,string(b1))
//
//	o2, e := f.Seek(6, 0)
//	check(e)
//	b2 := make([]byte,2)
//	n2, e := f.Read(b2)
//	check(e)
//	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))
//
//	o3, e := f.Seek(6, 0)
//	check(e)
//	b3:=make([]byte,2)
//	n3, e := io.ReadAtLeast(f, b3, 2)
//	check(e)
//	fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))
//
//	_, e = f.Seek(0, 0)
//	check(e)
//
//	r4 := bufio.NewReader(f)
//	b4, e := r4.Peek(5)
//	check(e)
//	fmt.Printf("5 bytes: %s\n",string(b4))
//
//	e = f.Close()
//
//}
