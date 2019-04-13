package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"html/template"
	"log"
	"mycode/trace"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	_ = t.templ.Execute(w, data)
}
func main() {
	var addr = flag.String("addr", ":8080", "the addr of the app.")
	flag.Parse()
	// setup gomniauth
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		github.New("af2b86ce2934d1e15b8d",
			"95e31723f3dac14c74184b66c370bc1335505661",
			"http://localhost:8080/auth/callback/github"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()
	log.Println("starting web server on", *addr)
	if e := http.ListenAndServe(*addr, nil); e != nil {
		log.Fatal("ListenAndServe", e)
	}
}
