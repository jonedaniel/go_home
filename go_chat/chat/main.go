package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	_ = t.templ.Execute(w, r)
}
func main() {
	var addr = flag.String("addr", ":8080", "the addr of the app.")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	log.Println("starting web server on", *addr)
	if e := http.ListenAndServe(*addr, nil); e != nil {
		log.Fatal("ListenAndServe", e)
	}
}
