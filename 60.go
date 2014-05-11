package main

import (
	"fmt"
    "net/http"
)

type String string

type Struct struct {
	Greeting string
    Punct string
    Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, s)
}

func main() {
    http.Handle("/string", String("Hello World!"))
    http.Handle("/struct", &Struct{"Welcome", ":", "johnnydiabetic"})
    http.ListenAndServe("localhost:4000", nil)
}