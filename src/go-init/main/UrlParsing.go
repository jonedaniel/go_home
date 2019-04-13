package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v1&k=v2#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	fmt.Println(u.Host)
	fmt.Println(strings.Split(u.Host, ":"))
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, err := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][1])

	param := "name=赵梦辉"
	sEnc := url.QueryEscape(param)
	fmt.Println(sEnc)
	dEnc, err := url.QueryUnescape(sEnc)
	fmt.Println(dEnc)

}
