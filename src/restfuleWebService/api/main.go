package main

import (
	"context"
	"flag"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)
func main() {
	var (
		HOST = os.Getenv("SP_FLEXIBLE_HOST")
		addr = flag.String("addr",":8080","endpoint address")
		mongo = flag.String("mongo",HOST,"mongodb address")
	)
	log.Println("Dialing mongo",*mongo)
	session, e := mgo.Dial(*mongo)
	if e != nil {
		log.Fatalln("fail to connect to mongo:",e)
	}
	defer session.Close()
	s := &Server{
		db: session,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/",withCORS(withAPIKey(s.handlePolls)))
	log.Println("Starting web server on", *addr)
	_ = http.ListenAndServe(":8080", mux)
	log.Println("Stopping")
}
type Server struct {
	db *mgo.Session
}

func withCORS(fn http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Expose-Headers","Location")
		w.Header().Set("Content-Type","application/json;charset=utf-8")
		fn(w,r)
	}
}
type contextKey struct {
	name string
}
var contextKeyAPIKey = &contextKey{"api-key"}

func APIKey(ctx context.Context)(string,bool){
	value := ctx.Value(contextKeyAPIKey)
	if value == nil {
		return "",false
	}
	valStr,ok := value.(string)
	return valStr,ok
}
func withAPIKey(fn http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAPIKey(key) {
			respondErr(w,r,http.StatusUnauthorized,"invalid api key")
			return
		}
		ctx := context.WithValue(r.Context(),contextKeyAPIKey,key)
		fn(w,r.WithContext(ctx))
	}
}
func isValidAPIKey(key string)bool{
	return key == "screw yo"
}